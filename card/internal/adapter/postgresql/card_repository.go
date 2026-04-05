package postgresql

import (
	"card/internal/consts"
	"card/internal/domain/entity"
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CardRepository struct {
	conn *pgxpool.Pool
}

func NewCardRepository(conn *pgxpool.Pool) *CardRepository {
	return &CardRepository{
		conn: conn,
	}
}

func (r *CardRepository) Create(ctx context.Context, model *entity.Card) error {
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

	_, err := r.conn.Exec(ctx, sql, args...)

	if err != nil {
		return fmt.Errorf("execute sql \"%s\": %w", sql, err)
	}

	return nil
}
