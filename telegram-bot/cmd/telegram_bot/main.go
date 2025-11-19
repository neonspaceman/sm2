package main

import (
	"context"
	"flag"
	"fmt"
	validatorPkg "github.com/go-playground/validator/v10"
	"github.com/go-telegram/bot"
	"os"
	"os/signal"
	dbalPkg "platform/pkg/dbal"
	logPkg "platform/pkg/logger"
	"sync"
	"syscall"
	"telegram-bot/internal/adapter/grpc"
	"telegram-bot/internal/adapter/postgresql"
	"telegram-bot/internal/api/bot_handler"
	"telegram-bot/internal/config"
	"telegram-bot/internal/consts"
	"telegram-bot/internal/usercase/command"
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

	dbal, err := dbalPkg.NewDBAL(cfg.Database.DSN, log)
	if err != nil {
		return err
	}
	defer dbal.Close()

	validator := validatorPkg.New(validatorPkg.WithRequiredStructEnabled())

	cardClient := grpc.NewCardClient()

	dialogRepository := postgresql.NewDialogRepository(dbal)
	userRepository := postgresql.NewUserRepository(dbal)

	dialogFirstOrCreateHandler := command.NewDialogFirstOrCreateHandler(dialogRepository)
	dialogUpdateHandler := command.NewDialogUpdateHandler(dialogRepository)
	userFirstOrCreateHandler := command.NewUserFirstOrCreateHandler(userRepository, validator)

	h := bot_handler.NewBotHandler(userFirstOrCreateHandler, dialogFirstOrCreateHandler, dialogUpdateHandler, cardClient, log)
	defer h.Close()

	b, err := bot.New(cfg.Telegram.BotToken, bot.WithDefaultHandler(h.Handle))

	if err != nil {
		panic(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	wg.Go(func() {
		fmt.Println("Telegram bot starting...")
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
