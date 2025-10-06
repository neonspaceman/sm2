package logger

import (
	"card/internal/config"
	"card/internal/consts"
	"card/pkg/request_id"
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zap *zap.Logger
}

func NewLogger(cfg *config.Config) (*Logger, error) {
	zapConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.Level(cfg.Log.Level)),
		Development:      cfg.App.Env == consts.EnvDev,
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

	return &Logger{zap: logger}, nil
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
