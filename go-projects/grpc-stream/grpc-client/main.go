package main

import (
	"context"
	api "github.com/billhcmus/grpc-stream/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:53001", grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("error connect to server, %v", err)
	}
	cli := api.NewPingPongClient(conn)
	pong, err := cli.Ping(context.Background(), &api.Ping{Timestamp: time.Now().Unix()})
	logrus.Infof("pong: %v", pong)
}
