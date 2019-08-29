package storage

import "sync/atomic"

type InMemoryOracle struct {
	ts int64
}

func NewInMemoryOracle() *InMemoryOracle {
	return &InMemoryOracle{0}
}

func (o *InMemoryOracle) getTS() int64 {
	return atomic.AddInt64(&o.ts, 1)
}