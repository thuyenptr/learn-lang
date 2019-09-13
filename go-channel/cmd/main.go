package main

import (
	"fmt"
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
