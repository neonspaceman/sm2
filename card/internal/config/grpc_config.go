package config

type GRPCConfig struct {
	Addr              string `envconfig:"APP_GRPC_ADDR"`
	MaxConnectionIdle int    `envconfig:"APP_GRPC_MAX_CONNECTION_IDLE"`
	Timeout           int    `envconfig:"APP_GRPC_TIMEOUT"`
	MaxConnectionAge  int    `envconfig:"APP_GRPC_MAX_CONNECTION_AGE"`
}
