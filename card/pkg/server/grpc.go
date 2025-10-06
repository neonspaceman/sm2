package grpc_server

import (
	apiPkg "card/internal/api/grpc"
	"card/internal/config"
	"card/internal/consts"
	"card/pkg/api/card"
	"card/pkg/dbal"
	"card/pkg/logger"
	"card/pkg/request_id"
	"context"
	"fmt"
	"github.com/google/uuid"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type GrpcServerProps struct {
	Cfg      *config.Config
	Log      *logger.Logger
	Dbal     *dbal.DBAL
	CardImpl *apiPkg.CardImpl
}

type GrpcServer struct {
	cfg      *config.Config
	log      *logger.Logger
	dbal     *dbal.DBAL
	cardImpl *apiPkg.CardImpl
}

func NewGrpcServer(props GrpcServerProps) *GrpcServer {
	return &GrpcServer{
		cfg:      props.Cfg,
		log:      props.Log,
		dbal:     props.Dbal,
		cardImpl: props.CardImpl,
	}
}

func (s *GrpcServer) Start() error {
	l, err := net.Listen("tcp", s.cfg.GRPC.Addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer l.Close()

	server := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Duration(s.cfg.GRPC.MaxConnectionIdle) * time.Minute,
			Timeout:           time.Duration(s.cfg.GRPC.Timeout) * time.Second,
			MaxConnectionAge:  time.Duration(s.cfg.GRPC.MaxConnectionAge) * time.Minute,
			Time:              time.Duration(s.cfg.GRPC.Timeout) * time.Minute,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			InterceptorRequestId(),
			grpc_logging.UnaryServerInterceptor(InterceptorLogger(s.log), grpc_logging.WithLogOnEvents(grpc_logging.StartCall, grpc_logging.FinishCall)),
		)),
	)

	card.RegisterCardServiceServer(server, s.cardImpl)

	if s.cfg.App.Env == consts.EnvDev {
		reflection.Register(server)
	}

	go func() {
		s.log.Info("GRPC server is listening...", zap.String("Addr", s.cfg.GRPC.Addr))
		if err = server.Serve(l); err != nil {
			s.log.Error("Failed running gRPC server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	server.GracefulStop()

	s.log.Info("GRPC server shut down correctly")

	return nil
}

func InterceptorLogger(l *logger.Logger) grpc_logging.Logger {
	return grpc_logging.LoggerFunc(func(ctx context.Context, lvl grpc_logging.Level, msg string, fields ...any) {
		switch lvl {
		case grpc_logging.LevelDebug:
			l.DebugCtx(ctx, msg)
		case grpc_logging.LevelInfo:
			l.InfoCtx(ctx, msg)
		case grpc_logging.LevelWarn:
			l.WarnCtx(ctx, msg)
		case grpc_logging.LevelError:
			l.ErrorCtx(ctx, msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func InterceptorRequestId() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		requestId := metadata.ExtractIncoming(ctx).Get(consts.GrpcRequestIdKey)

		if requestId != "" {
			requestId += " "
		}

		requestId += uuid.NewString()

		ctx = request_id.CtxSet(ctx, requestId)

		return handler(ctx, req)
	}
}
