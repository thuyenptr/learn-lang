package main

import (
	"github.com/sirupsen/logrus"
	"time"
)

type Request struct {
	value int
}

func handleRequest(reqs []*Request) chan *Request {
	done := make(chan *Request, 10)

	go func() {
	SEND:
		for _, req := range reqs {
			select {
			case done <- req:
			case <-time.After(10 * time.Second):
				logrus.Info("too slow")
				break SEND
			}
		}
	}()

	return done
}

func createDummyRequest(num int) []*Request {
	reqs := make([]*Request, 0, num)

	for i := 0; i < num; i++ {
		req := &Request{value: i}
		reqs = append(reqs, req)
	}

	return reqs
}

func main() {
	reqs := createDummyRequest(10)
	done := handleRequest(reqs)

	logrus.Infof("len %d", len(done))
}
