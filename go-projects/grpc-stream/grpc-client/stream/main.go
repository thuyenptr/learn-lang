package main

import (
	"context"
	api "github.com/billhcmus/grpc-stream/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	conn, err := grpc.Dial(":53001", grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
	}

	client := api.NewMathClient(conn)
	stream, err := client.Max(context.Background())
	if err != nil {
		logrus.Errorf("error open stream %v", err)
	}

	var max int64
	ctx := stream.Context()
	done := make(chan bool, 1)

	go func() {
		for i := 1; i <= 10; i++ {
			num := int64(rand.Intn(i))
			req := &api.MaxRequest{
				Num: num,
			}
			if err := stream.Send(req); err != nil {
				logrus.Fatalf("can not send, %v", err)
			}

			logrus.Infof("send %v", num)
			time.Sleep(time.Millisecond * 200)
		}

		if err := stream.CloseSend(); err != nil {
			logrus.Error(err)
		}
	}()

	go func() {
		for {
			maxRes, err := stream.Recv()
			if err != nil {
				logrus.Error(err)
			}
			if err == io.EOF {
				logrus.Info("stream end")
				close(done)
				return
			}

			max = maxRes.Max
			logrus.Infof("new max is %v", max)
		}
	}()
	
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			logrus.Error(err)
		}
		close(done)
	}()

	<-done
	logrus.Infof("finished with max=%d", max)
}
