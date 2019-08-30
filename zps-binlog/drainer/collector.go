package drainer

import (
	"context"
	"github.com/billhcmus/zps-binlog/drainer/checkpoint"
	"sync"
)


type notifyResult struct {
	err error
	wg sync.WaitGroup
}

type Collector struct {
	merger *Merger
	pumps     map[string]*Pump
	cp checkpoint.CheckPoint
	errCh chan error
	// notifyChan notifies the new pump is coming
	notifyChan chan *notifyResult
}

func NewCollector(cpt checkpoint.CheckPoint) *Collector {
	return &Collector{
		pumps: make(map[string]*Pump),
		errCh:           make(chan error, 10),
		merger: NewMerger(cpt.TS()),
		cp: cpt,
		notifyChan: make(chan *notifyResult),
	}
}

func (c *Collector) Start(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		c.publishBinlogs(ctx)
		wg.Done()
	}()

	c.keepUpdatingStatus(ctx, c.updateStatus)

	wg.Wait()
}

func (c *Collector) publishBinlogs(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case mergeItem, ok := <- c.merger.Output():
			if !ok {
				return
			}

			item := mergeItem.(*binlogItem)
			c.syncBinlog(item)
		}
	}
}

func (c *Collector) keepUpdatingStatus(ctx context.Context, updateStatus func(ctx context.Context)) {
	updateStatus(ctx)
}

func (c *Collector) updateStatus(ctx context.Context) {
	c.updatePumpStatus(ctx)

	c.updateCollectStatus(false)
}

func (c *Collector) updatePumpStatus(ctx context.Context) {
	// get pump nodes from register

	// for each nodes and do
	c.handlePumpStatusUpdate(ctx)
}

func (c *Collector) updateCollectStatus(synced bool) {

}

func (c *Collector) handlePumpStatusUpdate(ctx context.Context) {
	commitTS := c.merger.GetLatestTS()
	p := NewPump("fakenode", "127.0.0.1:1234", 1234, commitTS, c.errCh)
	c.merger.AddSource(MergeSource{
		ID: "fake-node",
		Source: p.PullBinlog(ctx, commitTS),
	})

}

func (c *Collector) syncBinlog(item *binlogItem) {
	// check job DML, DDL

	// Call Syncer
}