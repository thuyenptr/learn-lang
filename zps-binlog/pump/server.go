package pump

import (
	"context"
	"github.com/billhcmus/zps-binlog/pkg/util"
	pb "github.com/billhcmus/zps-binlog/proto"
	"github.com/billhcmus/zps-binlog/pump/storage"
	"github.com/pingcap/tidb/store/tikv"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	notifyDrainerTimeout            = time.Second * 10
	serverInfoOutputInterval        = time.Second * 10
	gcInterval                      = time.Hour
	earlyAlertGC                    = 20 * time.Hour
	detectDrainerCheckpointInterval = 10 * time.Minute
	GlobalConfig                    *globalConfig

	getPdClientFn         = util.GetPdClient
	newTiKVLockResolverFn = tikv.NewLockResolver
	//newKVStoreFn          = kvstore.New
)

type Server struct {
	dataDir string

	storage   storage.Storage
	clusterID uint64
}

func NewServer() *Server {
	st, err := storage.NewAppend("/tmp/binlog")
	if err != nil {
		logrus.Fatal(err)
		return nil
	}
	return &Server{
		storage: st,
	}
}

func (s *Server) WriteBinlogTest(blog *pb.Binlog) {
	if err := s.storage.WriteBinLog(blog); err != nil {
		return
	}
}

func (s *Server) writeBinlog(ctx context.Context, in *pb.WriteBinlogReq, isFakeBinlog bool) (*pb.WriteBinlogResp, error) {
	var err error
	blog := new(pb.Binlog)

	err = blog.XXX_Unmarshal(in.Payload)
	if err != nil {
		goto errHandle
	}

errHandle:
	return nil, err
}
