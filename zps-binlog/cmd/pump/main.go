package pump

import (
	pb "github.com/billhcmus/zps-binlog/proto"
	"github.com/billhcmus/zps-binlog/pump"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func init() {
	f, _ := os.OpenFile("log/zps-binlog.log", os.O_RDWR | os.O_CREATE, 0664)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(f)
}

func main() {
	//append, err := storage.NewAppend("/tmp/leveldb-zps")
	//if err != nil {
	//	logrus.Fatal(err)
	//}

	//testC := make(chan  int, 1024)
	//quit := make(chan struct{})
	//go func() {
	//	defer close(quit)
	//	for {
	//		select {
	//		case num, ok := <-testC:
	//			if !ok {
	//				logrus.Warn("not ok ", ok)
	//			}
	//			logrus.Info("get num ", num)
	//		}
	//	}
	//}()
	server := pump.NewServer()

	Pblog := &pb.Binlog{
		Tp: pb.BinlogType_Prewrite,
		StartTs: 1234,
	}

	Cblog := &pb.Binlog{
		Tp: pb.BinlogType_Commit,
		StartTs: 1234,
		CommitTs: 1239,
	}

	server.WriteBinlogTest(Pblog)

	time.Sleep(5 * time.Second)

	server.WriteBinlogTest(Cblog)
	select {

	}
}