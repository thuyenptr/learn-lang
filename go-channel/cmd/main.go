package main

import (
	"fmt"
	"github.com/billhcmus/go-channel/merger"
	"github.com/sirupsen/logrus"
	"sync"
)

/**
 * Created by thuyenpt
 * Date: 9/13/19
 */

type Handler struct {
	requestChan chan *Request
}

func NewHandler() *Handler {
	return &Handler{requestChan: make(chan *Request, 10)}
}

type Request struct {
	value int
	wg    sync.WaitGroup
}

func (h *Handler) convertValueToRequest(value int) {
	rq := &Request{
		value: value,
	}
	fmt.Printf("convert value %d\n", value)
	rq.wg.Add(1)
	h.requestChan <- rq
	rq.wg.Wait()
}

func (h *Handler) consumeRequest() {
	for req := range h.requestChan {
		fmt.Printf("request value %d\n", req.value)
		req.wg.Done()
	}
}

func main() {
	handler := NewHandler()
	go handler.consumeRequest()
	for i := 1; i < 5; i++ {
		handler.convertValueToRequest(i)
	}
}

func main1() {
	merge := merger.NewHeapStrategy()

	i1 := &merger.Item{
		Value: 10,
	}
	i2 := &merger.Item{
		Value: 1,
	}
	i3 := &merger.Item{
		Value: -1,
	}
	i4 := &merger.Item{
		Value: 12,
	}

	merge.Push(i1)
	merge.Push(i2)
	merge.Push(i3)
	merge.Push(i4)

	logrus.Info(merge.Pop())
}
