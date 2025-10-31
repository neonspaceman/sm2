package config

type TelegramConfig struct {
	BotToken string `envconfig:"APP_TELEGRAM_BOT_TOKEN"`
}
