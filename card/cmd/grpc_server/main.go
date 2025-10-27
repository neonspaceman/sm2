package main

import (
	"card/internal/adapter/postgresql"
	apiPkg "card/internal/api/grpc"
	"card/internal/config"
	"card/internal/consts"
	"card/internal/usecase/command"
	grpc_server "card/pkg/server"
	"flag"
	validatorPkg "github.com/go-playground/validator/v10"
	dbalPkg "platform/pkg/dbal"
	logPkg "platform/pkg/logger"
	"slices"
)

const defaultEnvFile = ".env"

func main() {
	var overrideEnv string
	flag.StringVar(&overrideEnv, "env", "", "path to override .env file")
	flag.Parse()

	envs := []string{defaultEnvFile}

	if overrideEnv != "" {
		envs = slices.Insert(envs, 0, overrideEnv)
	}

	cfg, err := config.Load(envs...)

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

	dbal, err := dbalPkg.NewDBAL(cfg, log)
	if err != nil {
		return err
	}
	defer dbal.Close()

	validator := validatorPkg.New(validatorPkg.WithRequiredStructEnabled())

	cardRepository := postgresql.NewCardRepository(dbal)

	cardCreateHandler := command.NewCardCreateHandler(cardRepository, validator)

	api := apiPkg.NewCardImpl(apiPkg.CardImplProps{
		CardCreateHandler: cardCreateHandler,
		Log:               log,
	})

	server := grpc_server.NewGrpcServer(grpc_server.GrpcServerProps{
		Cfg:      cfg,
		Log:      log,
		Dbal:     dbal,
		CardImpl: api,
	})
	return server.Start()
}
