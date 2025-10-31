package config

type LogConfig struct {
	Level int8 `envconfig:"APP_LOG_LEVEL"`
}
