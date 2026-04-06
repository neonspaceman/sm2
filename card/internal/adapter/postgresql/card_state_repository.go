package postgresql

import (
	"card/internal/consts"
	"card/internal/domain/entity"
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CardStateRepository struct {
	Repository
}

func NewCardStateRepository(pool *pgxpool.Pool, trm *trmpgx.CtxGetter) *CardStateRepository {
	return &CardStateRepository{
		Repository{
			pool: pool,
			trm:  trm,
		},
	}
}

func (r *CardStateRepository) Create(ctx context.Context, model *entity.CardState) error {
	sql, args := sq.
		Insert(consts.CardStateTableName).
		Columns(
			consts.CardStateIdColumn,
			consts.CardStateStateColumn,
			consts.CardStateStepColumn,
			consts.CardStateEasyColumn,
			consts.CardStateDueColumn,
			consts.CardStateCurrentIntervalInDaysColumn,
			consts.CardStateCreatedAtColumn,
			consts.CardStateUpdatedAtColumn,
		).
		Values(
			model.Id,
			model.State,
			model.Step,
			model.Easy,
			model.Due,
			model.CurrentIntervalInDays,
			model.CreatedAt,
			model.UpdatedAt,
		).
		PlaceholderFormat(sq.Dollar).
		MustSql()

	_, err := r.conn(ctx).Exec(ctx, sql, args...)

	if err != nil {
		return fmt.Errorf("execute sql \"%s\": %w", sql, err)
	}

	return nil
}
