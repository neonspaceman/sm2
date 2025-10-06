package config

type DatabaseConfig struct {
	DSN string `envconfig:"APP_DATABASE_DSN"`
}
