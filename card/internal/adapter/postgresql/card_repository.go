package postgresql

import (
	"card/internal/consts"
	card_domain "card/internal/domain/card"
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
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

func (r *CardRepository) Create(ctx context.Context, model *card_domain.Card) error {
	sql, args := sq.
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
		MustSql()

	_, err := r.conn(ctx).Exec(ctx, sql, args...)

	if err != nil {
		return fmt.Errorf("execute sql \"%s\": %w", sql, err)
	}

	return nil
}
