package main

func main() {
	//f, err := os.OpenFile("tikv.log", os.O_WRONLY | os.O_CREATE, 0755)
	//if err != nil {
	//	os.Exit(-1)
	//}
	//logrus.SetOutput(f)
	//if err := server.RunServer(); err != nil {
	//	logrus.Error(err)
	//	os.Exit(-1)
	//}

	maxCommitTSKey := []byte("!binlog!maxCommitTS")

	var maxCommitTS int64 = uint64(maxCommitTSKey)
}


