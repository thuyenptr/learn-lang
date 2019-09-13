package merger

import "container/heap"

/**
 * Created by thuyenpt
 * Date: 9/13/19
 */

type MergeItem interface {
	GetValue() int64
}

type MergeItems []MergeItem

func (m MergeItems) Len() int {
	return len(m)
}

func (m MergeItems) Less(i, j int) bool {
	return m[i].GetValue() < m[j].GetValue()
}

func (m MergeItems) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// implement heap interface

func (m *MergeItems) Push(x interface{}) {
	*m = append(*m, x.(MergeItem))
}

func (m *MergeItems) Pop() interface{} {
	old := *m
	top := old[len(old)-1]
	*m = old[0:(len(old) - 1)]
	return top
}

type HeapStrategy struct {
	items *MergeItems
}

func NewHeapStrategy() *HeapStrategy {
	h := &HeapStrategy{
		items: new(MergeItems),
	}
	heap.Init(h.items)
	return h
}

func (h *HeapStrategy) Push(item MergeItem) {
	heap.Push(h.items, item)
}

func (h *HeapStrategy) Pop() (item MergeItem) {
	if h.items.Len() == 0 {
		return nil
	}

	item = heap.Pop(h.items).(MergeItem)
	return
}

// item need sort
type Item struct {
	Value int64
}

func (i *Item) GetValue() int64 {
	return i.Value
}
