package grpc

import (
	"context"
	"github.com/billhcmus/tikv/pkg/api/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
)

func RunServer(ctx context.Context, v1API api.TiKVServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":" + port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	api.RegisterTiKVServiceServer(server, v1API)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			logrus.Info("shutting down gRPC server ...")
			server.GracefulStop()
			<- ctx.Done()
		}
	}()

	logrus.Info("gRPC server is running on port ", port)
	return server.Serve(listen)
}
