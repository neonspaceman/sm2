package postgresql

import (
	"card/internal/consts"
	card_domain "card/internal/domain/card"
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

type CardRepository struct {
	Repository
}

func NewCardRepository(pool *pgxpool.Pool, trm *trmpgx.CtxGetter) *CardRepository {
	return &CardRepository{
		Repository{
			pool: pool,
			trm:  trm,
		},
	}
}

func (r *CardRepository) GetById(ctx context.Context, id uuid.UUID) (*card_domain.Card, error) {
	sql, args, err := query_builder.CardQueryBuilder().
		Where(sq.Eq{consts.CardIdColumn: id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	rows, err := r.conn(ctx).Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("select card '%s': %w", id.String(), err)
	}

	model, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[card_domain.Card])

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, card_domain.NewCardNotFoundError(id)
	}

	if err != nil {
		return nil, fmt.Errorf("collect card '%s': %w", id.String(), err)
	}

	return model, nil
}

func (r *CardRepository) Create(ctx context.Context, model *card_domain.Card) error {
	sql, args, err := sq.
		Insert(consts.CardTableName).
		Columns(
			consts.CardIdColumn,
			consts.CardUserIdColumn,
			consts.CardQuestionColumn,
			consts.CardAnswerColumn,
			consts.CardFileTypeColumn,
			consts.CardFileIdColumn,
			consts.CardCreatedAtColumn,
			consts.CardUpdatedAtColumn,
		).
		Values(
			model.Id,
			model.UserId,
			model.Question,
			model.Answer,
			model.FileType,
			model.FileId,
			model.CreatedAt,
			model.UpdatedAt,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("build insert query: %w", err)
	}

	_, err = r.conn(ctx).Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("insert card: %w", err)
	}

	return nil
}
