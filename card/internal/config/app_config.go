package config

type AppConfig struct {
	Env string `envconfig:"APP_ENV"`
}
