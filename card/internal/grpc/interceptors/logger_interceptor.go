package interceptors

import (
	"context"
	"fmt"
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"platform/pkg/logger"
)

func LoggerInterceptor(l *logger.Logger) grpc_logging.Logger {
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
