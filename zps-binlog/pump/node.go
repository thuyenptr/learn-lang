package pump

import (
	"github.com/billhcmus/zps-binlog/pkg/etcd"
	"github.com/billhcmus/zps-binlog/pkg/file"
	"github.com/billhcmus/zps-binlog/pkg/flags"
	"github.com/billhcmus/zps-binlog/pkg/node"
	"github.com/juju/errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	shortIDLen = 8
	nodeIDFile = ".node"
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
	if err := checkExclusive(cfg.DataDir); err != nil {
		return nil, errors.Trace(err)
	}

	urlValues, err := flags.NewURLsValue(cfg.EtcdURLs)
	if err != nil {
		return nil, errors.Trace(err)
	}

	etcdCli, err := etcd.NewClientFromCfg(urlValues.StringSlice(), cfg.EtcdDialTimeout, node.DefaultRootPath, cfg.tls)
	if err != nil {
		return nil, errors.Trace(err)
	}

}

func generateLocalNodeID(dataDir string, listenAddr string) (string, error) {
	if err := os.MkdirAll(dataDir, file.PrivateDirMode); err != nil {
		return "", errors.Trace(err)
	}
}

// checkExclusive tries to get filelock of dataDir in exclusive mode
// if get lock fails, maybe some other pump is running
func checkExclusive(dataDir string) error {
	err := os.MkdirAll(dataDir, file.PrivateDirMode)
	if err != nil {
		return errors.Trace(err)
	}
	lockPath := filepath.Join(dataDir, lockFile)
	// when the process exits, the lockfile will be closed by system
	// and automatically release the lock
	_, err = file.TryLockFile(lockPath, os.O_WRONLY|os.O_CREATE, file.PrivateFileMode)
	return errors.Trace(err)
}

// FormatNodeID formats the nodeID, the nodeID should looks like "host:port"
func FormatNodeID(nodeID string) string {
	newNodeID := strings.TrimSpace(nodeID)

	return newNodeID
}
