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
	sql, args, err := query_builder.CardStateQueryBuilder().
		Where(sq.Eq{consts.CardStateIdColumn: id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	rows, err := r.conn(ctx).Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("select card '%s': %w", id.String(), err)
	}

	model, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[card_state_domain.CardState])

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, card_state_domain.ErrCardStateNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("collect card '%s': %w", id.String(), err)
	}

	return model, nil
}

func (r *CardStateRepository) Create(ctx context.Context, model *card_state_domain.CardState) error {
	sql, args, err := sq.
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
		ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	_, err = r.conn(ctx).Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("insert card: %w", err)
	}

	return nil
}

func (r *CardStateRepository) Save(ctx context.Context, model *card_state_domain.CardState) error {
	model.BeforeUpdate()

	sql, args, err := sq.
		Update(consts.CardStateTableName).
		Set(consts.CardStateStateColumn, model.State).
		Set(consts.CardStateStepColumn, model.Step).
		Set(consts.CardStateEasyColumn, model.Easy).
		Set(consts.CardStateDueColumn, model.Due).
		Set(consts.CardStateCurrentIntervalInDaysColumn, model.CurrentIntervalInDays).
		Set(consts.CardStateUpdatedAtColumn, model.UpdatedAt).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	_, err = r.conn(ctx).Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("update card: %w", err)
	}

	return nil
}
