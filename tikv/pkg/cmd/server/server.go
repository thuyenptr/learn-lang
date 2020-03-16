package server

import (
	"github.com/billhcmus/tikv/pkg/raftkv"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/raft"
	"go.etcd.io/etcd/raft/raftpb"
	"time"
)

type Config struct {
	gRPCPort string
}

func RunServer() error {
	//ctx := context.Background()

	node1 := raftkv.NewRaftNode(1, []raft.Peer{{ID: 1}, {ID: 2}, {ID: 3}})
	raftkv.AddNewRaftNode(node1)
	raftkv.StartCampaign(node1)
	go node1.Run()

	node2 := raftkv.NewRaftNode(2, []raft.Peer{{ID: 1}, {ID: 2}, {ID: 3}})
	raftkv.AddNewRaftNode(node2)
	go node2.Run()

	node3 := raftkv.NewRaftNode(3, []raft.Peer{{ID: 3}})
	raftkv.AddNewRaftNode(node3)
	go node3.Run()

	if err := raftkv.MakeProposeConfChange(2, raftpb.ConfChange{
		ID: 3,
		Type: raftpb.ConfChangeAddNode,
		NodeID: 3,
		Context: []byte(""),
	}); err != nil {
		logrus.Error(err.Error())
	}

	for node1.Raft.Status().Lead != 1 {
		time.Sleep(100 * time.Millisecond)
	}

	//raftkv.ShowPersistentStorage()

	//var conf Config
	//flag.StringVar(&conf.gRPCPort,"grpc-port", "", "gRPC port to bind")
	//flag.Parse()
	//v1API := v1.NewTestService()
	//return grpc.RunServer(ctx, v1API, conf.gRPCPort)
	return nil
}
