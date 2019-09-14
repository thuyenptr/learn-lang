package main

import (
	"fmt"
	"github.com/billhcmus/go-channel/merger"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

/**
 * Created by thuyenpt
 * Date: 9/13/19
 */

type Handler struct {
	requestChan chan *Request
}

func NewHandler() *Handler {
	return &Handler{
		requestChan: make(chan *Request, 10),
	}
}

type Request struct {
	value int
	wg    sync.WaitGroup
}

func (h *Handler) convertValueToRequest(value int) error {
	rq := &Request{
		value: value,
	}

	fmt.Printf("convert value %d\n", value)
	rq.wg.Add(1)
	h.requestChan <- rq
	rq.wg.Wait()
	return nil
}

func releaseRequest(batch []*Request) {
	fmt.Println("release request")
	for _, req := range batch {
		req.wg.Done()
	}
}

func (h *Handler) consumeRequest(wg sync.WaitGroup) {
	var count int
	var bufReqs []*Request

	for {
		select {
		case req, _ := <-h.requestChan:
			count++
			bufReqs = append(bufReqs, req)
			if count == 2 {
				releaseRequest(bufReqs)
				count = 0
				bufReqs = bufReqs[:0]
			}
			if count == 4 {
				wg.Done()
				//return
			}
			fmt.Printf("request value %d\n", req.value)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	handler := NewHandler()
	wg.Add(1)
	go handler.consumeRequest(wg)

	for i := 1; i < 5; i++ {
		_ = handler.convertValueToRequest(i)
	}

	time.Sleep(1 * time.Second)
	wg.Wait()
}
