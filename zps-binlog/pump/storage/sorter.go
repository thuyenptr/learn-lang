package storage

import (
	"container/list"
	pb "github.com/billhcmus/zps-binlog/proto"
	"github.com/sirupsen/logrus"
	"sync"
	"sync/atomic"
	"time"
)

type sortItem struct {
	start  int64
	commit int64
	tp     pb.BinlogType
}

type sorter struct {
	maxTSItemCB func(item sortItem)
	// nếu resolver trả về true, ta có thể bỏ qua P binlog, không cần phải đợi C binlog
	resolver func(startTs int64) bool
	// chứa startTS của txn ứng với C binlog cần đợi
	waitStartTS map[int64]struct{}

	lock sync.Mutex
	cond *sync.Cond
	items *list.List
	wg sync.WaitGroup
	closed int32
}

func newSorter(fn func(item sortItem)) *sorter {
	sorter := &sorter{
		maxTSItemCB: fn,
		items: list.New(),
		waitStartTS: make(map[int64]struct{}),
	}

	sorter.cond = sync.NewCond(&sorter.lock)
	//sorter.wg.Add(1)
	go sorter.run()
	return sorter
}

func (s *sorter) pushTSItem(item sortItem) {
	logrus.Info("pushTSItem push item to sorter")
	s.lock.Lock()

	if item.tp == pb.BinlogType_Prewrite {
		s.waitStartTS[item.start] = struct{}{}

		logrus.Infof("pushTSItem wait queue info %v", s.waitStartTS)
	} else {
		logrus.Info("pushTSItem receive binlog type commit")
		logrus.Infof("pushTSItem wait queue info %v", s.waitStartTS)
		delete(s.waitStartTS, item.start)
		logrus.Infof("pushTSItem after delete wait queue info %v", s.waitStartTS)
		logrus.Info("pushTSItem delete item from waitStartTS and call signal")
		s.cond.Signal()
	}

	s.items.PushBack(item)
	if s.items.Len() == 1 {
		logrus.Info("pushTSItem call signal")
		s.cond.Signal()
	}
	s.lock.Unlock()
}

func (s *sorter) run() {
	//defer s.wg.Done()
	go func() {
		tick := time.NewTicker(time.Second)

		for range tick.C {
			logrus.Info("run: tick in sorter, call signal")
			s.cond.Signal()
			if s.isClose() {
				return
			}
		}
	}()

	var maxTSItem sortItem
	for {
		s.cond.L.Lock()
		for s.items.Len() == 0 {
			if s.isClose() {
				s.cond.L.Unlock()
				return
			}
			logrus.Info("run: sorter wait...")
			s.cond.Wait()
			logrus.Info("run: get signal")
		}

		front := s.items.Front()
		item := front.Value.(sortItem)
		s.items.Remove(front)

		if item.tp == pb.BinlogType_Prewrite {
			logrus.Info("run: get prewrite binlog, let wait for commit binlog")
			getTime := time.Now()

			for {
				logrus.Info("run: wait Commit TS")
				_, ok := s.waitStartTS[item.start]
				if !ok {
					break
				}

				if time.Since(getTime) > time.Second {
					logrus.Info("run: do resolve")
					// TODO resolve here
					break
				}

				if s.isClose() {
					s.cond.L.Unlock()
					return
				}
				s.cond.Wait()
			}
		} else {
			logrus.Info("run: get commit binlog, let check this")
			if item.commit > maxTSItem.commit {
				logrus.Info("run: compare with maxTSItem...")
				maxTSItem = item
				s.maxTSItemCB(maxTSItem)
			}
		}
		s.cond.L.Unlock()
	}
}

func (s *sorter) isClose() bool {
	return atomic.LoadInt32(&s.closed) == 1
}

func (s *sorter) close() {
	atomic.StoreInt32(&s.closed, 1)

	s.lock.Lock()
	s.cond.Broadcast()
	s.lock.Unlock()
}



