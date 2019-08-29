package pump

import (
	"context"
	pb "github.com/billhcmus/zps-binlog/proto"
	"github.com/billhcmus/zps-binlog/pump/storage"
	"github.com/sirupsen/logrus"
)

type Server struct {
	storage storage.Storage
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