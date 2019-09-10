package pump

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/billhcmus/zps-binlog/pkg/security"
	"github.com/billhcmus/zps-binlog/pkg/util"
	"github.com/billhcmus/zps-binlog/pump/storage"
	"os"
	"time"
)

const (
	defaultEtcdDialTimeout   = 5 * time.Second
	defaultEtcdURLs          = "http://127.0.0.1:2379"
	dafaultListenAddr        = "127.0.0.1:8250"
	defaultMaxKafkaSize      = 1 << 30
	defaultHeartbeatInterval = 2
	defaultGC                = 7
	defaultDataDir           = "data.pump"

	// default interval time to generate fake binlog, the unit is second
	defaultGenFakeBinlogInterval = 3
)

// globalConfig is global config of pump to be sued in any where
type globalConfig struct {
	// enable online debug log output
	enableDebug bool
	// max binlog message size limit
	maxMsgSize int
}

// Config holds the configuration of pump
type Config struct {
	*flag.FlagSet
	LogLevel          string `toml:"log-level" json:"log-level"`
	NodeID            string `toml:"node-id" json:"node-id"`
	ListenAddr        string `toml:"addr" json:"addr"`
	AdvertiseAddr     string `toml:"advertise-addr" json:"advertise-addr"`
	Socket            string `toml:"socket" json:"socket"`
	EtcdURLs          string `toml:"pd-urls" json:"pd-urls"`
	EtcdDialTimeout   time.Duration
	DataDir           string `toml:"data-dir" json:"data-dir"`
	HeartbeatInterval int    `toml:"heartbeat-interval" json:"heartbeat-interval"`
	// pump only stores binlog events whose ts >= current time - GC(day)
	GC       int             `toml:"gc" json:"gc"`
	LogFile  string          `toml:"log-file" json:"log-file"`
	Security security.Config `toml:"security" json:"security"`

	GenFakeBinlogInterval int `toml:"gen-binlog-interval" json:"gen-binlog-interval"`

	MetricsAddr     string
	MetricsInterval int
	configFile      string
	printVersion    bool
	tls             *tls.Config
	Storage         storage.Config `toml:"storage" json:"storage"`
}

// NewConfig return an instance of configuration
func NewConfig() *Config {
	cfg := &Config{
		EtcdDialTimeout: defaultEtcdDialTimeout,
	}

	cfg.FlagSet = flag.NewFlagSet("pump", flag.ContinueOnError)
	fs := cfg.FlagSet
	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage of pump:")
		fs.PrintDefaults()
	}

	fs.StringVar(&cfg.NodeID, "node-id", "", "the ID of pump node; if not specified, we will generate one from hostname and the listening port")
	fs.StringVar(&cfg.ListenAddr, "addr", util.DefaultListenAddr(8250), "addr(i.e. 'host:port') to listen on for client traffic")
	fs.StringVar(&cfg.AdvertiseAddr, "advertise-addr", "", "addr(i.e. 'host:port') to advertise to the public")
	fs.StringVar(&cfg.Socket, "socket", "", "unix socket addr to listen on for client traffic")
	fs.StringVar(&cfg.EtcdURLs, "pd-urls", defaultEtcdURLs, "a comma separated list of the PD endpoints")
	fs.StringVar(&cfg.DataDir, "data-dir", "", "the path to store binlog data")
	fs.IntVar(&cfg.HeartbeatInterval, "heartbeat-interval", defaultHeartbeatInterval, "number of seconds between heartbeat ticks")
	fs.IntVar(&cfg.GC, "gc", defaultGC, "recycle binlog files older than gc days")
	fs.StringVar(&cfg.LogLevel, "L", "info", "log level: debug, info, warn, error, fatal")
	fs.StringVar(&cfg.MetricsAddr, "metrics-addr", "", "prometheus pushgateway address, leaves it empty will disable prometheus push")
	fs.IntVar(&cfg.MetricsInterval, "metrics-interval", 15, "prometheus client push interval in second, set \"0\" to disable prometheus push")
	fs.StringVar(&cfg.configFile, "config", "", "path to the pump configuration file")
	fs.BoolVar(&cfg.printVersion, "V", false, "print version information and exit")
	fs.StringVar(&cfg.LogFile, "log-file", "", "log file path")
	fs.IntVar(&cfg.GenFakeBinlogInterval, "fake-binlog-interval", defaultGenFakeBinlogInterval, "interval time to generate fake binlog, the unit is second")

	// global config
	fs.BoolVar(&GlobalConfig.enableDebug, "enable-debug", false, "enable print debug log")
	fs.IntVar(&GlobalConfig.maxMsgSize, "max-message-size", defaultMaxKafkaSize, "max msg size producer produce into kafka")
	fs.Int64Var(new(int64), "binlog-file-size", 0, "DEPRECATED")
	fs.BoolVar(new(bool), "enable-binlog-slice", false, "DEPRECATED")
	fs.IntVar(new(int), "binlog-slice-size", 0, "DEPRECATED")
	fs.StringVar(new(string), "log-rotate", "", "DEPRECATED")

	return cfg
}
