package main

import (
	"context"
	"flag"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"os"
	"os/signal"
	logPkg "platform/pkg/logger"
	"sync"
	"syscall"
	"telegram-bot/internal/config"
	"telegram-bot/internal/consts"
)

func main() {
	var envFile string
	flag.StringVar(&envFile, "env", "", "path to override .env file")
	flag.Parse()

	cfg, err := config.Load(envFile)

	if err != nil {
		panic(err)
	}

	err = run(cfg)

	if err != nil {
		panic(err)
	}
}

func run(cfg *config.Config) error {
	log, err := logPkg.NewLogger(
		logPkg.WithDevelopmentMode(cfg.App.Env == consts.EnvDev),
		logPkg.WithLevel(cfg.Log.Level),
	)

	if err != nil {
		return err
	}
	defer log.Close()

	b, err := bot.New(cfg.Telegram.BotToken, bot.WithDefaultHandler(func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		// ...
	}))

	if err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	wg.Go(func() {
		log.Info("Telegram bot starting...")
		b.Start(ctx)
	})

	// Wait SIGTERM
	<-quit

	// Cancel context
	cancel()

	// Wait graceful shutdown
	wg.Wait()

	return nil
}
