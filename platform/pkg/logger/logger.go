package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"platform/pkg/request_id"
)

type Logger struct {
	level       int8
	development bool

	zap *zap.Logger
}

func NewLogger(opts ...Option) (*Logger, error) {
	l := &Logger{}

	for _, opt := range opts {
		opt(l)
	}

	zapConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.Level(l.level)),
		Development:      l.development,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	logger = logger.WithOptions(zap.AddCallerSkip(1))

	l.zap = logger

	return l, nil
}

func (log *Logger) Debug(msg string, fields ...zap.Field) {
	log.zap.Debug(msg, fields...)
}

func (log *Logger) Info(msg string, fields ...zap.Field) {
	log.zap.Info(msg, fields...)
}

func (log *Logger) Warn(msg string, fields ...zap.Field) {
	log.zap.Warn(msg, fields...)
}

func (log *Logger) Error(msg string, fields ...zap.Field) {
	log.zap.Error(msg, fields...)
}

func (log *Logger) DebugCtx(ctx context.Context, msg string, fields ...zap.Field) {
	log.zap.Debug(msg, append(fields, log.requestId(ctx))...)
}

func (log *Logger) InfoCtx(ctx context.Context, msg string, fields ...zap.Field) {
	log.zap.Info(msg, append(fields, log.requestId(ctx))...)
}

func (log *Logger) WarnCtx(ctx context.Context, msg string, fields ...zap.Field) {
	log.zap.Warn(msg, append(fields, log.requestId(ctx))...)
}

func (log *Logger) ErrorCtx(ctx context.Context, msg string, fields ...zap.Field) {
	log.zap.Error(msg, append(fields, log.requestId(ctx))...)
}

func (log *Logger) Close() {
	log.zap.Sync()
}

func (log *Logger) requestId(ctx context.Context) zap.Field {
	value := request_id.CtxGet(ctx)

	if value == "" {
		return zap.Skip()
	}

	return zap.String("request_id", value)
}
