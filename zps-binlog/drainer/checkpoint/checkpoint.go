package checkpoint


type CheckPoint interface {
	TS() int64
}

func NewCheckPoint() CheckPoint {
	return newMySqlCheckpoint()
}