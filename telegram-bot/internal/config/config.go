package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Log      LogConfig
	Telegram TelegramConfig
	GRPC     GRCPConfig
}

func Load(path ...string) (*Config, error) {
	err := godotenv.Load(path...)
	if err != nil {
		return nil, err
	}

	cfg := Config{}

	err = envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
