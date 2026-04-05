package config

type GRCPConfig struct {
	CardURI string `envconfig:"APP_GRPC_CARD_URI"`
}
