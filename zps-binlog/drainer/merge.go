package drainer

import "sync"

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