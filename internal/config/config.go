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
	Endpoint string `env:"RUN_ADDRESS" envDefault:":8082"`
	AppName  string `env:"APP_NAME" envDefault:"CaptchaMaker"`
	Debug    bool   `env:"CAP_MAKER_APP_SERVER_DEBUG"`
}

// InitConfig initialises config, first from flags, then from env, so that env overwrites flags
func InitConfig() (*Config, error) {
	var cfg Config
	flag.StringVar(&cfg.Endpoint, "a", "127.0.0.1:8081", "server address as host:port")
	flag.BoolVar(&cfg.Debug, "debug", true, "key for hash function")
	flag.Parse()

	err := env.Parse(&cfg)

	if err != nil {
		return nil, err
	}
	return &cfg, nil
}