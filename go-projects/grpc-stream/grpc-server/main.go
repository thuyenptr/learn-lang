package main

import (
	"context"
	api "github.com/billhcmus/grpc-stream/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"net"
	"time"
)

type server struct {}

func NewPingPongService() *server {
	return &server{}
}

func (s *server) Ping(ctx context.Context, ping *api.Ping) (*api.Pong, error) {
	logrus.Infof("ping: %v", ping)
	return &api.Pong{
		Timestamp: time.Now().Unix(),
	}, nil
}

type serverstream struct {}

func (s *serverstream) Max(srv api.Math_MaxServer) error {
	logrus.Info("server stream start and serve Max rpc")
	var max int64
	ctx := srv.Context()

	for {
		select {
		case <- ctx.Done():
			return ctx.Err()
		default:
		}
		req, err := srv.Recv()
		if err == io.EOF {
			logrus.Info("close stream")
			return nil
		}
		if err != nil {
			logrus.Warnf("receive err %v", err)
		}

		if req.Num <= max {
			continue
		}

		max = req.Num
		res := &api.MaxResponse{
			Max: max,
		}

		if err := srv.Send(res); err != nil {
			logrus.Warn("send error %v", err)
		}
		logrus.Infof("send success max = %v", max)
	}
}

func NewServerStream() *serverstream {
	return &serverstream{}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:53001")
	if err != nil {
		logrus.Errorf("error listening, %v", err)
	}
	server := grpc.NewServer()
	api.RegisterMathServer(server, NewServerStream())
	logrus.Infof("grpc server listening at %v", 53001)
	if err := server.Serve(lis); err != nil {
		logrus.Errorf("error serve, %v", err)
	}
}
