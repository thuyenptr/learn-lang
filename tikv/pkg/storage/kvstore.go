package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tecbot/gorocksdb"
)

//type KVStorage struct {
//}

const (
	DbPath ="/tmp/gorocksdb"
)

func OpenDB() (*gorocksdb.DB, []*gorocksdb.ColumnFamilyHandle, error) {
	options := gorocksdb.NewDefaultOptions()
	options.SetCreateIfMissingColumnFamilies(true)
	options.SetCreateIfMissing(true)

	//db, err := gorocksdb.OpenDb(options, DbPath)
	//if err != nil {
	//	return nil, err
	//}

	cfNames := []string{"default", "lock", "write"}
	cfOpts := []*gorocksdb.Options{
		gorocksdb.NewDefaultOptions(),
		gorocksdb.NewDefaultOptions(),
		gorocksdb.NewDefaultOptions()}

	db, cfhandles, err := gorocksdb.OpenDbColumnFamilies(options, DbPath, cfNames, cfOpts)
	if err != nil {
		return nil, nil, err
	}
	return db, cfhandles, nil
}

func main() {
	db, handles, err := OpenDB()
	if err != nil {
		logrus.Fatal(err)
		return
	}

	defer func() {
		for _, handle := range handles {
			handle.Destroy()
		}
	}()

	readOptions := gorocksdb.NewDefaultReadOptions()
	readOptions.SetFillCache(true)

	writeOptions := gorocksdb.NewDefaultWriteOptions()
	writeOptions.SetSync(true)

	//for i := 0; i < 100; i++ {
	//	keyStr := "zlp" + strconv.Itoa(i)
	//	key := []byte(keyStr)
	//	if err := db.PutCF(writeOptions, handles[0], key, key); err != nil {
	//		logrus.Warn("Put key error, ", err)
	//	}
	//}
	//
	//logrus.Infof("Put complete")
	//
	//for i := 0; i < 100; i++ {
	//	keyStr := "zlp" + strconv.Itoa(i)
	//	key := []byte(keyStr)
	//	//data, err := db.Get(readOptions, key)
	//	data, err := db.GetCF(readOptions, handles[0], key)
	//	if err != nil {
	//		logrus.Warn("Get error, ", err)
	//	}
	//
	//	logrus.Infof("Data size: %v, value: %v", data.Size(), string(data.Data()))
	//}

	batch := gorocksdb.NewWriteBatch()
	batch.PutCF(handles[1], []byte("hello"), []byte("word"))
	batch.PutCF(handles[0], []byte("hi"), []byte("word"))
	if err := db.Write(writeOptions, batch); err != nil {
		logrus.Error(err)
	}

	value, _  := db.GetCF(readOptions, handles[1], []byte("hello"))
	logrus.Info("value: ", string(value.Data()))
}