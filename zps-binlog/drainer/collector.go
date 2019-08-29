package drainer

import (
	"context"
	"github.com/billhcmus/zps-binlog/drainer/checkpoint"
	"sync"
)

type Collector struct {
	merger *Merger
	cp checkpoint.CheckPoint
}

func NewCollector(cpt checkpoint.CheckPoint) *Collector {
	merger := NewMerger(cpt.TS())
	return &Collector{
		merger: merger,
		cp: cpt,
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
	//var commitTS int64 = 0
	//c.merger.AddSource(MergeSource {ID: "dummy", Source: p.PullBinlog(ctx, commitTS)})
}

func (c *Collector) syncBinlog(item *binlogItem) {
	// check job DML, DDL

	// Call Syncer
}