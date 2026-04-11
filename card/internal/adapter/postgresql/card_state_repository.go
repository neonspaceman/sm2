package postgresql

import (
	"card/internal/consts"
	card_state_domain "card/internal/domain/card_state"
	"card/internal/query_builder"
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

func (r *CardStateRepository) GetById(ctx context.Context, id uuid.UUID) (*card_state_domain.CardState, error) {
	sql, args := query_builder.CardStateQueryBuilder().
		Where(sq.Eq{consts.CardStateIdColumn: id}).
		MustSql()

	rows, err := r.conn(ctx).Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("query_builder: %w", err)
	}

	model, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[card_state_domain.CardState])

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, card_state_domain.ErrCardStateNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("collect one: %w", err)
	}

	return model, nil
}

func (r *CardStateRepository) Create(ctx context.Context, model *card_state_domain.CardState) error {
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
