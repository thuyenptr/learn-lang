package drainer

import pb "github.com/billhcmus/zps-binlog/proto"

type binlogItem struct {
	binlog *pb.Binlog
	nodeID string
}


func (b *binlogItem) GetCommitTs() int64 {
	return b.binlog.CommitTs
}

func (b *binlogItem) GetSourceID() string {
	return b.nodeID
}

func newBinlogItem(b *pb.Binlog, nodeID string) *binlogItem {
	return &binlogItem{
		binlog: b,
		nodeID: nodeID,
	}
}