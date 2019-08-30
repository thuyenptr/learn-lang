package drainer

import (
	"context"
	"encoding/json"
	pb "github.com/billhcmus/zps-binlog/proto"
	"github.com/pingcap/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"sync/atomic"
	"time"
)

const (
	binlogChanSize = 10
)

type Pump struct {
	nodeID string
	addr string
	clusterID uint64

	// the latest binlog ts that pump had handled
	latestTS int64

	isClosed int32
	isPaused int32

	errCh chan error

	pullCli pb.Pump_PullBinlogsClient
	grpcConn *grpc.ClientConn
}

func NewPump(nodeID, addr string, clusterID uint64, startTS int64, errCh chan error) *Pump {
	nodeID = strings.TrimSpace(nodeID)
	return &Pump{
		nodeID: nodeID,
		addr: addr,
		clusterID: clusterID,
		latestTS: startTS,
		errCh: errCh,
	}
}

func (p *Pump) Close() {
	logrus.Info("pump is closing")
	atomic.StoreInt32(&p.isClosed, 1)
}

func (p *Pump) Pause() {
	if atomic.CompareAndSwapInt32(&p.isPaused, 0, 1) {
		logrus.Info("pump pause pull binlog")
	}
}

func (p *Pump) Continue(ctx context.Context) {
	if atomic.CompareAndSwapInt32(&p.isPaused, 1, 0) {
		logrus.Info("pump continue pull binlog")
	}
}

func (p *Pump) PullBinlog(ctx context.Context, last int64) (ret chan MergeItem) {
	ret = make(chan MergeItem, 10)

	go func() {
		logrus.Info("pump start PullBinlog")

		needReCreateConn := false
		for {
			if atomic.LoadInt32(&p.isClosed) == 1 {
				return
			}

			if atomic.LoadInt32(&p.isPaused) == 1 {
				// this pump is paused, wait until it can pull binlog again
				logrus.Info("pump is paused")
				time.Sleep(time.Second)
				continue
			}

			if p.grpcConn == nil || needReCreateConn {
				logrus.Info("pump create pull binlogs client")
				if err := p.createPullBinlogsClient(ctx, last); err != nil {
					logrus.Error("pump create pull binlogs client failed")
					time.Sleep(time.Second)
					continue
				}

				needReCreateConn = false
			}

			resp, err := p.pullCli.Recv()
			if err != nil {
				if status.Code(err) != codes.Canceled {
					logrus.Error("pump receive binlog failed")
				}

				needReCreateConn = true

				time.Sleep(time.Second)
				// TODO: add metric here
				continue
			}

			binlog := new(pb.Binlog)
			err = json.Unmarshal(resp.Entity.Payload, binlog)
			if err != nil {
				logrus.Error("pump unmarshal binlog failed")
				return
			}

			item := newBinlogItem(binlog, p.nodeID)

			select {
			case ret <- item:
				if binlog.CommitTs > last {
					last = binlog.CommitTs
					p.latestTS = binlog.CommitTs
				} else {
					logrus.Error("pump receive un-sort binlog")
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return
}

func (p *Pump) createPullBinlogsClient(ctx context.Context, last int64) error {
	if p.grpcConn != nil {
		_ = p.grpcConn.Close()
	}

	callOpts := []grpc.CallOption{grpc.MaxCallRecvMsgSize(1 << 30)}


	conn, err := grpc.Dial(p.addr, grpc.WithInsecure(), grpc.WithDefaultCallOptions(callOpts...))
	if err != nil {
		logrus.Error("pump create grpc dial failed")
		p.pullCli = nil
		p.grpcConn = nil
		return errors.Trace(err)
	}

	cli := pb.NewPumpClient(conn)

	in := &pb.PullBinlogReq{
		ClusterID: p.clusterID,
		StartFrom: &pb.Pos{Offset: last},
	}

	pullCli, err := cli.PullBinlogs(ctx, in)
	if err != nil {
		logrus.Error("pump create PullBinlogs client failed")
		_ = conn.Close()
		p.pullCli = nil
		p.grpcConn = nil
		return errors.Trace(err)
	}

	p.pullCli = pullCli
	p.grpcConn = conn

	return nil
}