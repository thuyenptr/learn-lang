package raftkv

import (
	"bytes"
	"context"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/raft"
	"go.etcd.io/etcd/raft/raftpb"
	"math"
	"os"
	"time"
)

type raftNode struct {
	// node id in the cluster
	id uint64
	// implement by etc-d, node in the cluster will use
	Raft raft.Node
	// raftkv configuration
	conf *raft.Config
	// context
	ctx context.Context
	// store log entries, state in-memory
	store *raft.MemoryStorage
	// fake persistent store
	kvStore map[string]string

	ticker <- chan time.Time

	done <- chan struct{}
}

var nodes = make(map[uint64]*raftNode)

func NewRaftNode(id uint64, peers []raft.Peer) *raftNode {
	hb := 1 // heart beat
	store := raft.NewMemoryStorage()
	raftConfig := &raft.Config {
		// ID in etc-d Raft-kv
		ID: uint64(id),
		// ElectionTick timeout before election occur
		ElectionTick: hb * 10,
		// HeartbeatTick time between hearts beat
		HeartbeatTick: hb,
		// Storage for raftkv to store entries and state
		Storage: store,
		// Max byte size for each message
		MaxSizePerMsg: math.MaxUint16,
		// Max number of message in-flight during optimistic replication phase
		MaxInflightMsgs: 256,
	}

	node := &raftNode{
		id: id,
		conf: raftConfig,
		store: store,
		kvStore: make(map[string]string),
		ticker: time.Tick(time.Second),
		ctx: context.TODO(),
	}
	node.Raft = raft.StartNode(raftConfig, peers)
	return node
}

func AddNewRaftNode(node *raftNode) {
	nodes[node.id] = node
}

func StartCampaign(node *raftNode) {
	if err := node.Raft.Campaign(node.ctx); err != nil {
		logrus.Error(err.Error())
		return
	}
}

func ShowPersistentStorage() {
	for i, node := range nodes {
		logrus.Infof("Show persistent of node: %v", i)
		for k, v := range node.kvStore {
			logrus.Infof("key: %v - value: %v", k, v)
		}
		logrus.Info("******************")
	}
}

func MakeProposeConfChange(nodeID uint64, confChange raftpb.ConfChange) error {
	if err := nodes[nodeID].Raft.ProposeConfChange(nodes[nodeID].ctx, confChange); err != nil {
		return err
	}
	return nil
}

func (n *raftNode)Run() {
	for {
		select {
		case <-n.ticker:
			n.Raft.Tick()
		case rd := <- n.Raft.Ready():
			if err := n.saveToStorage(rd.HardState, rd.Entries, rd.Snapshot); err != nil {
				logrus.Error(err.Error())
				os.Exit(-1)
			}
			if err := n.send(rd.Messages); err != nil {
				logrus.Error(err.Error())
				os.Exit(-1)
			}

			for _, entry := range rd.CommittedEntries {
				n.process(entry)
				if entry.Type == raftpb.EntryConfChange {
					var confChange raftpb.ConfChange
					if err := confChange.Unmarshal(entry.Data); err != nil {
						logrus.Error(err.Error())
					}
					n.Raft.ApplyConfChange(confChange)
				}
			}
			n.Raft.Advance()
		case <- n.done:
			return
		}
	}
}

func (n *raftNode)saveToStorage(hardState raftpb.HardState, entries []raftpb.Entry, snapshot raftpb.Snapshot) error {
	if err := n.store.Append(entries); err != nil {
		return err
	}
	if !raft.IsEmptyHardState(hardState) {
		if err := n.store.SetHardState(hardState); err != nil {
			return err
		}
	}
	if !raft.IsEmptySnap(snapshot) {
		if err := n.store.ApplySnapshot(snapshot); err != nil {
			return err
		}
	}
	return nil
}

func (n *raftNode)send(message []raftpb.Message) error {
	for _,msg := range message {
		logrus.Infof("Message detail: %v", raft.DescribeMessage(msg, nil))
		// send message to other node
		if err := nodes[msg.To].receive(n.ctx, msg); err != nil {
			return err
		}
	}
	return nil
}

func (n *raftNode)receive(ctx context.Context, message raftpb.Message) error {
	return n.Raft.Step(ctx, message)
}

func (n *raftNode)process(entry raftpb.Entry) {
	logrus.Infof("node %v, process entry %v\n", n.id, entry)
	logrus.Infof("entry type %v", entry.Type)
	if entry.Type == raftpb.EntryNormal && entry.Data != nil {
		parts := bytes.SplitN(entry.Data, []byte(":"), 2)
		n.kvStore[string(parts[0])] = string(parts[1])
	}
}