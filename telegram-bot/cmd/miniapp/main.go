package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	validator_pkg "github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	log_pkg "platform/pkg/logger"
	"reflect"
	"strings"
	"telegram-bot/internal/adapter/grpc"
	"telegram-bot/internal/adapter/postgresql"
	"telegram-bot/internal/api/miniapp/rest"
	"telegram-bot/internal/config"
	"telegram-bot/internal/consts"
	"telegram-bot/internal/http/middleware"
	"telegram-bot/internal/usercase/command"
	"telegram-bot/internal/usercase/query"
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
	log, err := log_pkg.NewLogger(
		log_pkg.WithDevelopmentMode(cfg.App.Env == consts.EnvDev),
		log_pkg.WithLevel(cfg.Log.Level),
	)

	if err != nil {
		return err
	}
	defer log.Close()

	conn, err := pgxpool.New(context.Background(), cfg.Database.DSN)
	if err != nil {
		return err
	}
	err = conn.Ping(context.Background())
	if err != nil {
		return err
	}
	defer conn.Close()

	validator := validator_pkg.New(validator_pkg.WithRequiredStructEnabled())
	validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		if field.Tag.Get("json") != "" {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name != "-" {
				return name
			}
		}

		if field.Tag.Get("form") != "" {
			return field.Tag.Get("form")
		}

		return ""
	})

	cardClient := grpc.NewCardClient(cfg)

	userRepository := postgresql.NewUserRepository(conn)

	getCardsHandler := query.NewGetCardByUserIdHandler(cardClient)
	userFirstOrCreateHandler := command.NewUserFirstOrCreateHandler(userRepository)

	// TODO: remove
	botToken := "5768337691:AAH5YkoiEuPk8-FZa32hStHTqXiLPtAEhx8"

	r := gin.Default()
	r.Use(middleware.TelegramAuth(botToken), middleware.ErrorHandler(log), gin.Recovery())

	cardHandler := rest.NewCardHandler(userFirstOrCreateHandler, getCardsHandler, validator)
	cardHandler.RegisterRoutes(r)

	// Start miniapp on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	return r.Run()
}
