package main

import (
	"card/internal/adapter/postgresql"
	api_pkg "card/internal/api/grpc"
	"card/internal/config"
	"card/internal/consts"
	grpc_pkg "card/internal/grpc/interceptors"
	"card/internal/usecase/command"
	"card/internal/usecase/query"
	"card/pkg/api/card"
	"context"
	"flag"
	"fmt"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	validator_pkg "github.com/go-playground/validator/v10"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	log_pkg "platform/pkg/logger"
	"slices"
	"syscall"
	"time"
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

	cfg, errors := config.Load(envs...)

	if errors != nil {
		panic(errors)
	}

	errors = run(cfg)

	if errors != nil {
		panic(errors)
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

	pool, err := pgxpool.New(context.Background(), cfg.Database.DSN)
	if err != nil {
		return err
	}
	err = pool.Ping(context.Background())
	if err != nil {
		return err
	}
	defer pool.Close()

	validator := validator_pkg.New(validator_pkg.WithRequiredStructEnabled())
	trManager, err := manager.New(trmpgx.NewDefaultFactory(pool))

	if err != nil {
		return err
	}

	cardRepository := postgresql.NewCardRepository(pool, trmpgx.DefaultCtxGetter)
	cardStateRepository := postgresql.NewCardStateRepository(pool, trmpgx.DefaultCtxGetter)

	createCardHandler := command.NewCreateCardHandler(cardRepository, cardStateRepository, trManager, validator)
	getCardsByUserIdHandler := query.NewGetCardByUserIdHandler(pool, validator)

	api := api_pkg.NewCardImpl(api_pkg.CardImplProps{
		CreateCardHandler:     createCardHandler,
		GetCardsByUserIdQuery: getCardsByUserIdHandler,
		Log:                   log,
	})

	// GRPC Server
	l, err := net.Listen("tcp", cfg.GRPC.Addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer l.Close()

	server := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Duration(cfg.GRPC.MaxConnectionIdle) * time.Minute,
			Timeout:           time.Duration(cfg.GRPC.Timeout) * time.Second,
			MaxConnectionAge:  time.Duration(cfg.GRPC.MaxConnectionAge) * time.Minute,
			Time:              time.Duration(cfg.GRPC.Timeout) * time.Minute,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			//InterceptorRequestId(),
			grpc_logging.UnaryServerInterceptor(grpc_pkg.LoggerInterceptor(log), grpc_logging.WithLogOnEvents(grpc_logging.StartCall, grpc_logging.FinishCall)),
		)),
	)

	card.RegisterCardServiceServer(server, api)

	if cfg.App.Env == consts.EnvDev {
		reflection.Register(server)
	}

	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		log.Info("GRPC server is listening...", zap.String("Addr", cfg.GRPC.Addr))
		err := server.Serve(l)

		if err != nil {
			return err
		}

		log.Info("GRPC server shut down correctly")

		return err
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	server.GracefulStop()

	return g.Wait()
}
