package pump

import (
	"context"
	"fmt"
	"github.com/billhcmus/zps-binlog/pkg/etcd"
	"github.com/billhcmus/zps-binlog/pkg/flags"
	"github.com/billhcmus/zps-binlog/pkg/node"
	"github.com/billhcmus/zps-binlog/pkg/util"
	pb "github.com/billhcmus/zps-binlog/proto"
	"github.com/juju/errors"
	"github.com/pingcap/log"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	shortIDLen = 8
	lockFile   = ".lock"
)

var nodePrefix = "pumps"

type pumpNode struct {
	sync.RWMutex
	*node.EtcdRegistry

	status            *node.Status
	heartbeatInterval time.Duration

	// latestTS and latestTime is used for get approach ts
	latestTS   int64
	latestTime time.Time

	// use this function to update max commit ts
	getMaxCommitTS func() int64
}

// NewPumpNode returns a pumpNode obj that initialized by server config
func NewPumpNode(cfg *Config, getMaxCommitTS func() int64) (node.Node, error) {
	//if err := checkExclusive(cfg.DataDir); err != nil {
	//	return nil, errors.Trace(err)
	//}

	urlValues, err := flags.NewURLsValue(cfg.EtcdURLs)
	if err != nil {
		return nil, errors.Trace(err)
	}

	etcdCli, err := etcd.NewClientFromCfg(urlValues.StringSlice(), cfg.EtcdDialTimeout, node.DefaultRootPath, cfg.tls)
	if err != nil {
		return nil, errors.Trace(err)
	}

	nodeID, err := readNodeID()
	if err != nil {
		if cfg.NodeID != "" {
			nodeID = cfg.NodeID
		} else if errors.IsNotFound(err) {
			nodeID, err = generateNodeID(cfg.DataDir, cfg.ListenAddr)
			if err != nil {
				return nil, errors.Trace(err)
			}
		} else {
			return nil, errors.Trace(err)
		}
		return nil, errors.Trace(err)
	} else if cfg.NodeID != "" {
		log.Warn("you had a node ID in local file.[if you want to change the node ID, you should delete the file data-dir/.node file]")
	}

	advURL, err := url.Parse(cfg.AdvertiseAddr)
	if err != nil {
		return nil, errors.Annotatef(err, "invalid configuration of advertise addr(%s)", cfg.AdvertiseAddr)
	}

	status := &node.Status{
		NodeID:  nodeID,
		Addr:    advURL.Host,
		State:   node.Paused,
		IsAlive: true,
	}

	return &pumpNode{
		EtcdRegistry:      node.NewEtcdRegistry(etcdCli, cfg.EtcdDialTimeout),
		status:            status,
		heartbeatInterval: time.Duration(cfg.HeartbeatInterval) * time.Second,
		getMaxCommitTS:    getMaxCommitTS,
	}, nil
}

func (p *pumpNode) ID() string {
	return p.status.NodeID
}

func (p *pumpNode) Close() error {
	return errors.Trace(p.EtcdRegistry.Close())
}

func (p *pumpNode) ShortID() string {
	if len(p.status.NodeID) <= shortIDLen {
		return p.status.NodeID
	}
	return p.status.NodeID[0:shortIDLen]
}

func (p *pumpNode) RefreshStatus(ctx context.Context, status *node.Status) error {
	p.Lock()
	defer p.Unlock()

	p.status = status
	if p.status.UpdateTS != 0 {
		p.latestTS = p.status.UpdateTS
		p.latestTime = time.Now()
	} else {
		p.updateStatus()
	}

	err := p.UpdateNode(ctx, nodePrefix, status)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (p *pumpNode) Notify(ctx context.Context) error {
	drainers, err := p.Nodes(ctx, "drainers")
	if err != nil {
		return errors.Trace(err)
	}

	dialerOpts := []grpc.DialOption{
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			if deadline, ok := ctx.Deadline(); ok {
				return net.DialTimeout("tcp", addr, time.Until(deadline))
			}
			return net.DialTimeout("tcp", addr, 0)
		}),

		grpc.WithInsecure(),
		grpc.WithBlock(),
	}

	for _, c := range drainers {
		if c.State != node.Online {
			continue
		}
		if err := notifyDrainer(ctx, c, dialerOpts); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}

func notifyDrainer(ctx context.Context, c *node.Status, dialerOpts []grpc.DialOption) error {
	log.Info("Start trying to notify drainer", zap.String("addr", c.Addr))
	var clientConn *grpc.ClientConn
	err := util.RetryContext(ctx, 3, time.Second, 2, func(ictx context.Context) error {
		log.Info("Connecting drainer", zap.String("addr", c.Addr))
		var err error
		clientConn, err = grpc.DialContext(ictx, c.Addr, dialerOpts...)
		return err
	})
	if err != nil {
		return errors.Annotatef(err, "connect drainer(%s)", c.Addr)
	}
	defer func() {
		_ = clientConn.Close()
	}()

	drainer := pb.NewCisternClient(clientConn)

	err = util.RetryContext(ctx, 3, time.Second, 2, func(ictx context.Context) error {
		log.Info("Notifying drainer", zap.String("addr", c.Addr))
		_, err := drainer.Notify(ictx, nil)
		return err
	})
	if err != nil {
		return errors.Annotatef(err, "notify drainer(%s)", c.Addr)
	}
	return nil
}

func (p *pumpNode) NodeStatus() *node.Status {
	return p.status
}

func (p *pumpNode) NodesStatus(ctx context.Context) ([]*node.Status, error) {
	return p.Nodes(ctx, nodePrefix)
}

func (p *pumpNode) Heartbeat(ctx context.Context) <-chan error {
	errC := make(chan error, 1)
	go func() {
		defer func() {
			close(errC)
			log.Info("Heartbeat goroutine exited")
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(p.heartbeatInterval):
				p.Lock()
				p.updateStatus()
				if err := p.UpdateNode(ctx, nodePrefix, p.status); err != nil {
					errC <- errors.Trace(err)
				}
				p.Unlock()
			}
		}
	}()
	return errC
}

func (p *pumpNode) Quit() error {
	return errors.Trace(p.Close())
}

func (p *pumpNode) updateStatus() {
	p.status.UpdateTS = util.GetApproachTS(p.latestTS, p.latestTime)
	p.status.MaxCommitTS = p.getMaxCommitTS()
}

func generateNodeID(dataDir string, listenAddr string) (string, error) {
	urlLis, err := url.Parse(listenAddr)
	if err != nil {
		return "", errors.Trace(err)
	}
	_, port, err := net.SplitHostPort(urlLis.Host)
	if err != nil {
		return "", errors.Trace(err)
	}
	hostName, err := os.Hostname()
	if err != nil {
		return "", errors.Trace(err)
	}
	nodeID := FormatNodeID(fmt.Sprintf("%s:%s", hostName, port))

	// TODO store nodeID to etcd
	return nodeID, nil
}

func readNodeID() (string, error) {
	// TODO get nodeID from etcd

	return "", nil
}

// checkExclusive tries to get filelock of dataDir in exclusive mode
// if get lock fails, maybe some other pump is running
//func checkExclusive(dataDir string) error {
//	err := os.MkdirAll(dataDir, file.PrivateDirMode)
//	if err != nil {
//		return errors.Trace(err)
//	}
//	lockPath := filepath.Join(dataDir, lockFile)
//	// when the process exits, the lockfile will be closed by system
//	// and automatically release the lock
//	_, err = file.TryLockFile(lockPath, os.O_WRONLY|os.O_CREATE, file.PrivateFileMode)
//	return errors.Trace(err)
//}

// FormatNodeID formats the nodeID, the nodeID should looks like "host:port"
func FormatNodeID(nodeID string) string {
	newNodeID := strings.TrimSpace(nodeID)

	return newNodeID
}
