package storage

import (
	"context"
	"encoding/binary"
	pb "github.com/billhcmus/zps-binlog/proto"
	"github.com/pingcap/errors"
	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
	"go.uber.org/zap"
	"sync"
	"time"
)

const (
	chanCapacity = 1 << 20
)

var (
	// save gcTS, the max TS we have gc, for binlog not greater than gcTS, we can delete it from storage
	gcTSKey = []byte("!binlog!gcts")
	// save maxCommitTS, we can get binlog in range [gcTS, maxCommitTS]  from PullCommitBinlog
	maxCommitTSKey = []byte("!binlog!maxCommitTS")
	// save the valuePointer we should start push binlog item to sorter when restart
	handlePointerKey = []byte("!binlog!handlePointer")
	// save valuePointer headPointer, for binlog in vlog not after headPointer, we have save it in metadata db
	// at start up, we can scan the vlog from headPointer and save the ts -> valuePointer to metadata db
	headPointerKey = []byte("!binlog!headPointer")
	// If the kv channel blocks for more than this value, turn on the slow chaser
	slowChaserThreshold = 3 * time.Second
	// Values of handlePointer and MaxCommitTS will be saved at most once per interval
	handlePtrSaveInterval = time.Second
)

type Storage interface {
	WriteBinLog(binlog *pb.Binlog) error

	// delete <= ts
	GC(ts int64)
	GetGCTS() int64

	// AllMatched return if all the P-binlog have the matching C-binlog
	AllMatched() bool

	MaxCommitTS() int64

	// GetBinlog return the binlog of ts
	GetBinlog(ts int64) (binlog *pb.Binlog, err error)
	// PullCommitBinlog return the chan to consume the binlog
	PullCommitBinlog(ctx context.Context, last int64) <-chan []byte

	Close() error
}

type Append struct {
	dir      string
	metadata *leveldb.DB
	sorter   *sorter
	writeCh  chan *request
	latestTs int64

	headPointer   valuePointer
	handlePointer valuePointer

	gcTS        int64
	maxCommitTS int64

	sortItems          chan sortItem
	handleSortItemQuit chan struct{}

	wg sync.WaitGroup
}

func NewAppend(dir string) (append *Append, err error) {
	metadata, err := openMetadataDB(dir)
	if err != nil {
		return nil, errors.Trace(err)
	}

	writeCh := make(chan *request, chanCapacity)

	append = &Append{
		metadata:  metadata,
		writeCh:   writeCh,
		sortItems: make(chan sortItem, 1024),
	}

	append.gcTS, err = append.readGCTSFromDB()
	if err != nil {
		return nil, errors.Trace(err)
	}
	// TODO set metrics here

	append.maxCommitTS, err = append.readInt64(maxCommitTSKey)
	if err != nil {
		return nil, errors.Trace(err)
	}
	// TODO set metrics here

	append.headPointer, err = append.readPointer(headPointerKey)
	if err != nil {
		return nil, errors.Trace(err)
	}

	append.handlePointer, err = append.readPointer(handlePointerKey)
	if err != nil {
		return nil, errors.Trace(err)
	}

	append.handleSortItemQuit = append.handleSortItem(append.sortItems)
	sorter := newSorter(func(item sortItem) {
		logrus.Info("maxTSItemCB: sorter get item", item)
		append.sortItems <- item
	})

	append.sorter = sorter

	toKV := append.writeToValueLog(writeCh)

	append.wg.Add(1)
	go append.writeToSorter(append.writeToKV(toKV))

	append.wg.Add(1)
	go append.updateStatus()

	return
}

func (a *Append) readInt64(key []byte) (int64, error) {
	value, err := a.metadata.Get(key, nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return 0, nil
		}
		return 0, errors.Trace(err)
	}
	return int64(binary.LittleEndian.Uint64(value)), nil
}

func (a *Append) readGCTSFromDB() (int64, error) {
	return a.readInt64(gcTSKey)
}

func (a *Append) saveGCTSToDB(ts int64) error {
	value := make([]byte, 8)
	binary.LittleEndian.PutUint64(value, uint64(ts))
	return a.metadata.Put(gcTSKey, value, nil)
}

func (a *Append) readPointer(key []byte) (valuePointer, error) {
	var vp valuePointer
	value, err := a.metadata.Get(key, nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return vp, nil
		}
		return vp, errors.Trace(err)
	}

	err = vp.UnmarshalBinary(value)
	if err != nil {
		return vp, errors.Trace(err)
	}
	return vp, nil
}

//func (a *Append) Push() {
//	req := &request{
//
//	}
//}

func (a *Append) handleSortItem(items <-chan sortItem) (quit chan struct{}) {
	logrus.Info("handleSortItem invoke...")
	quit = make(chan struct{})
	go func() {
		defer close(quit)

		var toSaveItem sortItem
		var toSave <-chan time.Time

		for {
			select {
			case item, ok := <-items:
				logrus.Info("handleSortItem receive item")
				if !ok {
					if toSave != nil {
						err := a.persistHandlePointer(toSaveItem)
						if err != nil {
							logrus.Error(errors.ErrorStack(err))
						}
					}
				}

				toSaveItem = item
				if toSave == nil {
					toSave = time.After(handlePtrSaveInterval)
				}

				logrus.Info("handleSortItem get sort item ", zap.Reflect("item", item))

			case <-toSave:
				logrus.Info("handleSortItem persisting handle pointer")
				toSave = nil
			}
		}
	}()
	return quit
}

func (a *Append) persistHandlePointer(item sortItem) error {
	return nil
}

func (a *Append) Close() error {
	logrus.Info("Close close append")
	close(a.sortItems)
	<-a.handleSortItemQuit

	err := a.metadata.Close()
	if err != nil {
		logrus.Error("Close close metadata failed", zap.Error(err))
	}

	return err
}

func (a *Append) writeToValueLog(reqs chan *request) chan *request {
	done := make(chan *request, chanCapacity)
	go func() {
		for {
			logrus.Info("writeToValueLog loop for wait")
			// select blocking here
			select {
			case req, ok := <-reqs:
				logrus.Info("writeToValueLog get request from writeCh")
				if !ok {
					logrus.Info("writeToValueLog !ok return")
					return
				}
				done <- req
			}
		}
	}()
	return done
}

func (a *Append) writeToSorter(reqs chan *request) {
	defer a.wg.Done()
	for req := range reqs {
		logrus.Info("writeToSorter write request to sorter")
		var item sortItem
		item.start = req.startTS
		item.commit = req.commitTS
		item.tp = req.tp

		a.sorter.pushTSItem(item)
	}
}

func (a *Append) writeToKV(reqs chan *request) chan *request {
	done := make(chan *request, 128)

	go func() {
		defer close(done)
		for req := range reqs {
			logrus.Info("writeToKV write binlog to kv")
			done <- req
		}
	}()

	return done
}

func (a *Append) WriteBinLog(binlog *pb.Binlog) error {
	a.writeBinlog(binlog)
	return nil
}

func (a *Append) writeBinlog(binlog *pb.Binlog) {
	logrus.Info("writeBinlog write binlog receive binlog & create request")
	req := new(request)

	req.payload = nil
	req.startTS = binlog.StartTs
	req.commitTS = binlog.CommitTs
	req.tp = binlog.Tp

	//req.wg.Add(1)
	logrus.Info("writeBinlog send request to write channel")
	a.writeCh <- req
}

func (a *Append) GC(ts int64) {
	panic("implement me")
}

func (a *Append) MaxCommitTS() int64 {
	panic("implement me")
}

func (a *Append) GetBinlog(ts int64) (binlog *pb.Binlog, err error) {
	panic("implement me")
}

func (a *Append) PullCommitBinlog(ctx context.Context, last int64) <-chan []byte {
	panic("implement me")
}

func (a *Append) updateStatus() {
	defer a.wg.Done()

	updateLatestTicker := time.NewTicker(time.Second)
	updateLatest := updateLatestTicker.C

	for {
		select {
		case <-updateLatest:
			// TODO: get latest timestamp from PD
		}
	}
}
