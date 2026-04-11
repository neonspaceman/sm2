package postgresql

import (
	"card/internal/consts"
	review_domain "card/internal/domain/review"
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ReviewLogRepository struct {
	Repository
}

func NewReviewLogRepository(pool *pgxpool.Pool, trm *trmpgx.CtxGetter) *ReviewLogRepository {
	return &ReviewLogRepository{
		Repository{
			pool: pool,
			trm:  trm,
		},
	}
}

func (r *ReviewLogRepository) Create(ctx context.Context, model *review_domain.ReviewLog) error {
	sql, args, err := sq.
		Insert(consts.ReviewLogTableName).
		Columns(
			consts.ReviewLogIdColumn,
			consts.ReviewLogCardIdColumn,
			consts.ReviewLogRatingColumn,
			consts.ReviewLogCreatedAtColumn,
		).
		Values(
			model.Id,
			model.CardId,
			model.Rating,
			model.CreatedAt,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	_, err = r.conn(ctx).Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("insert review log: %w", err)
	}

	return nil
}
