package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
)

// InitLogger configures zap logger
func InitLogger(debug bool, projectID string) (*zap.Logger, error) {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.LevelKey = "severity"
	zapConfig.EncoderConfig.MessageKey = "message"

	if debug {
		zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	logger, err := zapConfig.Build(zap.Fields(
		zap.String("projectID", projectID),
	))
	if err != nil {
		return nil, err
	}

	return logger, nil
}

type Config struct {
	Endpoint   string `env:"RUN_ADDRESS" envDefault:":8082"`
	TimeOut500 string `env:"TIME_OUT" envDefault:"1"`
	SizeCache  string `env:"SIZE_CACHE" envDefault:"100000"`
	AppName    string `env:"APP_NAME" envDefault:"CaptchaMaker"`
	Debug      bool   `env:"CAP_MAKER_APP_SERVER_DEBUG"`
}

// InitConfig initialises config, first from flags, then from env, so that env overwrites flags
func InitConfig() (*Config, error) {
	var cfg Config
	flag.StringVar(&cfg.Endpoint, "a", ":8081", "server address as host:port")
	flag.StringVar(&cfg.TimeOut500, "t", "1", "timeout request")
	flag.StringVar(&cfg.SizeCache, "s", "100000", "size cache")

	flag.BoolVar(&cfg.Debug, "debug", true, "key for hash function")
	flag.Parse()

	err := env.Parse(&cfg)

	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
