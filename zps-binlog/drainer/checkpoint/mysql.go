package checkpoint


type MysqlCheckPoint struct {
	CommitTS int64
}

func (mc *MysqlCheckPoint) TS() int64 {
	return mc.CommitTS
}

func newMySqlCheckpoint() *MysqlCheckPoint {
	return &MysqlCheckPoint{}
}