package storage

import (
	"encoding/binary"
	"errors"
	pb "github.com/billhcmus/zps-binlog/proto"
	"sync"
)

type Options struct {
	ValueLogFileSize          int64
	Sync                      bool
	KVChanCapacity            int
	SlowWriteThreshold        float64
	StopWriteAtAvailableSpace uint64

	KVConfig *KVConfig
}

type request struct {
	startTS  int64
	commitTS int64
	tp       pb.BinlogType

	payload      []byte
	valuePointer valuePointer
	wg           sync.WaitGroup
	err          error
}

func (r *request) ts() int64 {
	if r.tp == pb.BinlogType_Prewrite {
		return r.startTS
	}

	return r.commitTS
}

type valuePointer struct {
	Fid    uint32
	Offset int64
}

func (vp *valuePointer) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 12)
	binary.LittleEndian.PutUint32(data, vp.Fid)
	binary.LittleEndian.PutUint64(data[4:], uint64(vp.Offset))
	return
}

func (vp *valuePointer) UnmarshalBinary(data []byte) error {
	if len(data) < 12 {
		return errors.New("not enough data")
	}
	vp.Fid = binary.LittleEndian.Uint32(data)
	vp.Offset = int64(binary.LittleEndian.Uint64(data[4:]))
	return nil
}

type valueLog struct {
}
