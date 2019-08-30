package drainer

import (
	"sync"
	"sync/atomic"
)

type MergeItem interface {
	GetCommitTs() int64
	GetSourceID() string
}

type MergeItems []MergeItem

func (m MergeItems) Len() int {
	return len(m)
}

func (m MergeItems) Less(i, j int) bool {
	return m[i].GetCommitTs() < m[j].GetCommitTs()
}

func (m MergeItems) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type MergeStrategy interface {
	Push(MergeItem)
	Pop() MergeItem
	Exist(string) bool
}

type Merger struct {
	sync.RWMutex

	sources map[string]MergeSource
	strategy MergeStrategy
	output chan MergeItem

	latestTS int64

	close int32
	pause int32
}

type MergeSource struct {
	ID string
	Source chan MergeItem
}

func NewMerger(ts int64, sources ...MergeSource) *Merger {
	m := &Merger{
		latestTS: ts,
		sources: make(map[string]MergeSource),
		output: make(chan MergeItem, 10),
	}

	for i := 0; i < len(sources); i++ {
		m.sources[sources[i].ID] = sources[i]
	}

	go m.run()
	return m
}

func (m *Merger) run() {

}

func (m *Merger) AddSource(source MergeSource) {
	m.Lock()
	m.sources[source.ID] = source
	// do something here
	m.Unlock()
}

func (m *Merger) Output() chan MergeItem {
	return m.output
}

// GetLatestTS returns the last binlog's ts send to syncer
func (m *Merger) GetLatestTS() int64 {
	m.RLock()
	defer m.RUnlock()
	return m.latestTS
}

// Stop stops merge
func (m *Merger) Stop() {
	atomic.StoreInt32(&m.pause, 1)
}

// Continue continue merge
func (m *Merger) Continue() {
	atomic.StoreInt32(&m.pause, 0)
}

func (m *Merger) isPaused() bool {
	return atomic.LoadInt32(&m.pause) == 1
}
