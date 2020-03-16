package main

import (
	"context"
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
	return &Handler{
		requestChan: make(chan *Request, 10),
	}
}

type Request struct {
	value int
	wg    sync.WaitGroup
}

func (h *Handler) convertValueToRequest(wg *sync.WaitGroup, value int) error {
	rq := &Request{
		value: value,
	}

	logrus.Infof("convert value %d\n", value)
	rq.wg.Add(1)
	h.requestChan <- rq
	rq.wg.Wait()
	wg.Done()
	return nil
}

func (h *Handler) displayChannelSize(ctx context.Context) {
	go func() {
		for {
			logrus.Infof("[Size] Channel size %d", len(h.requestChan))
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func releaseRequest(batch []*Request) {
	for _, req := range batch {
		logrus.Infof("release request value %d", req.value)
		req.wg.Done()
	}
}

func consumeRequest(ctx context.Context, reqs chan *Request) {
	var count int
	var bufReqs []*Request

	for {
		select {
		case req, ok := <-reqs:
			logrus.Infof("case receive request from channel")
			if !ok {
				return
			}

			count++
			bufReqs = append(bufReqs, req)
			if count > 2 {
				logrus.Info("count > 2")
				releaseRequest(bufReqs)
				count = 0
				bufReqs = bufReqs[:0]
			}
			logrus.Infof("request value %d\n", req.value)
		case <-ctx.Done():
			return
		default:
			logrus.Info("case default")
			if len(bufReqs) > 0 {
				releaseRequest(bufReqs)
				count = 0
				bufReqs = bufReqs[:0]
				continue
			}

			req, ok := <-reqs
			if !ok {
				return
			}
			bufReqs = append(bufReqs, req)
			count++
		}
	}
}

func main() {
	var wg sync.WaitGroup
	handler := NewHandler()
	ctx, cancelFunc := context.WithCancel(context.Background())
	go consumeRequest(ctx, handler.requestChan)
	go handler.displayChannelSize(ctx)

	for i := 1; i < 50; i++ {
		wg.Add(1)
		go func() {
			_ = handler.convertValueToRequest(&wg, i)
		}()
	}

	wg.Wait()
	cancelFunc()
	close(handler.requestChan)
}
