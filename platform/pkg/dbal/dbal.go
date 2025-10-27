package dbal

import (
	"card/internal/config"
	"card/pkg/logger"
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type DBAL struct {
	psql *pgxpool.Pool
	log  *logger.Logger
}

func NewDBAL(config *config.Config, log *logger.Logger) (*DBAL, error) {
	log.Info(fmt.Sprintf("Connection to %s", config.Database.DSN))

	connection, err := pgxpool.New(context.Background(), config.Database.DSN)
	if err != nil {
		return nil, err
	}

	err = connection.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	log.Info(fmt.Sprintf("Connected to %s", config.Database.DSN))

	return &DBAL{connection, log}, nil
}

func (dbal *DBAL) SqlBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

func (dbal *DBAL) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	dbal.log.DebugCtx(ctx, "DBAL query", zap.String("sql", sql), zap.Any("args", args))
	return dbal.psql.Query(ctx, sql, args...)
}

func (dbal *DBAL) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	dbal.log.DebugCtx(ctx, "DBAL exec", zap.String("sql", sql), zap.Any("args", args))
	return dbal.psql.Exec(ctx, sql, args...)
}

func (dbal *DBAL) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	dbal.log.DebugCtx(ctx, "DBAL query row", zap.String("sql", sql), zap.Any("args", args))
	return dbal.psql.QueryRow(ctx, sql, args...)
}

func (dbal *DBAL) ScanOne(ctx context.Context, dest any, sql string, args ...any) error {
	row, err := dbal.Query(ctx, sql, args...)
	defer row.Close()

	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, row)
}

func (dbal *DBAL) ScanAll(ctx context.Context, dest any, sql string, args ...any) error {
	row, err := dbal.Query(ctx, sql, args...)
	defer row.Close()

	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, row)
}

func (dbal *DBAL) Close() {
	dbal.psql.Close()
}
