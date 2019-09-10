package storage

import (
	"github.com/pingcap/log"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"go.uber.org/zap"
	"golang.org/x/sys/unix"
	"reflect"
)

// Config holds the configuration of the storage
type Config struct {
	SyncLog *bool `toml:"sync-log" json:"sync-log"`
	// the channel to buffer binlog meta, pump will block write binlog request if the channel is full
	KVChanCapacity            int            `toml:"kv_chan_cap" json:"kv_chan_cap"`
	SlowWriteThreshold        float64        `toml:"slow_write_threshold" json:"slow_write_threshold"`
	KV                        *KVConfig      `toml:"kv" json:"kv"`
	StopWriteAtAvailableSpace *HumanizeBytes `toml:"stop-write-at-available-space" json:"stop-write-at-available-space"`
}

// Configuration for storage layer of the pump
type KVConfig struct {
	BlockCacheCapacity            int     `toml:"block-cache-capacity" json:"block-cache-capacity"`
	BlockRestartInterval          int     `toml:"block-restart-interval" json:"block-restart-interval"`
	BlockSize                     int     `toml:"block-size" json:"block-size"`
	CompactionL0Trigger           int     `toml:"compaction-L0-trigger" json:"compaction-L0-trigger"`
	CompactionTableSize           int     `toml:"compaction-table-size" json:"compaction-table-size"`
	CompactionTotalSize           int     `toml:"compaction-total-size" json:"compaction-total-size"`
	CompactionTotalSizeMultiplier float64 `toml:"compaction-total-size-multiplier" json:"compaction-total-size-multiplier"`
	WriteBuffer                   int     `toml:"write-buffer" json:"write-buffer"`
	WriteL0PauseTrigger           int     `toml:"write-L0-pause-trigger" json:"write-L0-pause-trigger"`
	WriteL0SlowdownTrigger        int     `toml:"write-L0-slowdown-trigger" json:"write-L0-slowdown-trigger"`
}

var defaultStorageKVConfig = &KVConfig{
	BlockCacheCapacity:            8 * opt.MiB,
	BlockRestartInterval:          16,
	BlockSize:                     4 * opt.KiB,
	CompactionL0Trigger:           8,
	CompactionTableSize:           64 * opt.MiB,
	CompactionTotalSize:           512 * opt.MiB,
	CompactionTotalSizeMultiplier: 8,
	WriteBuffer:                   64 * opt.MiB,
	WriteL0PauseTrigger:           24,
	WriteL0SlowdownTrigger:        17,
}

func setDefaultStorageConfig(cf *KVConfig) {
	if cf.BlockCacheCapacity <= 0 {
		cf.BlockCacheCapacity = defaultStorageKVConfig.BlockCacheCapacity
	}
	if cf.BlockRestartInterval <= 0 {
		cf.BlockRestartInterval = defaultStorageKVConfig.BlockRestartInterval
	}
	if cf.BlockSize <= 0 {
		cf.BlockSize = defaultStorageKVConfig.BlockSize
	}
	if cf.CompactionL0Trigger <= 0 {
		cf.CompactionL0Trigger = defaultStorageKVConfig.CompactionL0Trigger
	}
	if cf.CompactionTableSize <= 0 {
		cf.CompactionTableSize = defaultStorageKVConfig.CompactionTableSize
	}
	if cf.CompactionTotalSize <= 0 {
		cf.CompactionTotalSize = defaultStorageKVConfig.CompactionTotalSize
	}
	if cf.CompactionTotalSizeMultiplier <= 0 {
		cf.CompactionTotalSizeMultiplier = defaultStorageKVConfig.CompactionTotalSizeMultiplier
	}
	if cf.WriteBuffer <= 0 {
		cf.WriteBuffer = defaultStorageKVConfig.WriteBuffer
	}
	if cf.WriteL0PauseTrigger <= 0 {
		cf.WriteL0PauseTrigger = defaultStorageKVConfig.WriteL0PauseTrigger
	}
	if cf.WriteL0SlowdownTrigger <= 0 {
		cf.WriteL0SlowdownTrigger = defaultStorageKVConfig.WriteL0SlowdownTrigger
	}
}

func openMetadataDB(kvDir string, cf *KVConfig) (*leveldb.DB, error) {
	if cf == nil {
		cf = defaultStorageKVConfig
	} else {
		setDefaultStorageConfig(cf)
	}

	log.Info("open metadata db", zap.Reflect("config", cf))

	var opts opt.Options
	opts.BlockCacheCapacity = cf.BlockCacheCapacity
	opts.BlockRestartInterval = cf.BlockRestartInterval
	opts.BlockSize = cf.BlockSize
	opts.CompactionL0Trigger = cf.CompactionL0Trigger
	opts.CompactionTableSize = cf.CompactionTableSize
	opts.CompactionTotalSize = cf.CompactionTotalSize
	opts.CompactionTotalSizeMultiplier = cf.CompactionTotalSizeMultiplier
	opts.WriteBuffer = cf.WriteBuffer
	opts.WriteL0PauseTrigger = cf.WriteL0PauseTrigger
	opts.WriteL0SlowdownTrigger = cf.WriteL0SlowdownTrigger

	return leveldb.OpenFile(kvDir, &opts)
}

type storageSize struct {
	capacity  uint64
	available uint64
}

func getStorageSize(dir string) (size storageSize, err error) {
	var stat unix.Statfs_t

	err = unix.Statfs(dir, &stat)
	if err != nil {
		return size, err
	}

	// When container is run in MacOS, `bsize` obtained by `statfs` syscall is not the fundamental block size,
	// but the `iosize` (optimal transfer block size) instead, it's usually 1024 times larger than the `bsize`.
	// for example `4096 * 1024`. To get the correct block size, we should use `frsize`. But `frsize` isn't
	// guaranteed to be supported everywhere, so we need to check whether it's supported before use it.
	// For more details, please refer to: https://github.com/docker/for-mac/issues/2136
	bSize := uint64(stat.Bsize)
	field := reflect.ValueOf(&stat).Elem().FieldByName("Frsize")
	if field.IsValid() {
		if field.Kind() == reflect.Uint64 {
			bSize = field.Uint()
		} else {
			bSize = uint64(field.Int())
		}
	}

	// Available blocks * size per block = available space in bytes
	size.available = stat.Bavail * bSize
	size.capacity = stat.Blocks * bSize
	return
}
