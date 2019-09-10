package storage

import (
	"encoding/binary"
	"github.com/dustin/go-humanize"
	"github.com/pingcap/errors"
	"sync/atomic"
)

var tsKeyPrefix = []byte("ts:")

func decodeTSKey(key []byte) int64 {
	// check bound
	_ = key[len(tsKeyPrefix)+8-1]
	return int64(binary.BigEndian.Uint64(key[len(tsKeyPrefix):]))
}

func encodeTSKey(ts int64) []byte {
	buf := make([]byte, 8+len(tsKeyPrefix))
	copy(buf, tsKeyPrefix)

	binary.BigEndian.PutUint64(buf[len(tsKeyPrefix):], uint64(ts))
	return buf
}

// HumanizeBytes is used for humanize configure
type HumanizeBytes int64

// Uint64 return bytes
func (b HumanizeBytes) Uint64() uint64 {
	return uint64(b)
}

func (b *HumanizeBytes) UnmarshalText(text []byte) (err error) {
	err = nil
	if len(text) == 0 {
		*b = 0
		return
	}
	num, err := humanize.ParseBytes(string(text))
	if err != nil {
		return errors.Annotatef(err, "text: %s", string(text))
	}
	*b = HumanizeBytes(num)
	return
}

// test helper
type memOracle struct {
	ts int64
}

func newMemOracle() *memOracle {
	return &memOracle{
		ts: 0,
	}
}

func (o *memOracle) getTS() int64 {
	return atomic.AddInt64(&o.ts, 1)
}
