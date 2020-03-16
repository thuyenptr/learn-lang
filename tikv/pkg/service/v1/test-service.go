package v1

import (
	"context"
	"github.com/billhcmus/tikv/pkg/api/v1"
	"time"
)

type testServiceServer struct {}

// NewTestService function create new testService instance
func NewTestService() *testServiceServer {
	return &testServiceServer{}
}

func (ts *testServiceServer) Ping(ctx context.Context, ping *api.Ping) (*api.Pong, error) {
	return &api.Pong{
		Timestamp: time.Now().String(),
		ServiceName: "Test Service",
	}, nil
}

