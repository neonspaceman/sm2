package postgresql

import (
	"card/internal/domain/entity"
	"context"
	"fmt"
	"platform/pkg/dbal"
)

const (
	cardTableName       = "card"
	cardIdColumn        = "id"
	cardUserIdColumn    = "user_id"
	cardQuestionColumn  = "question"
	cardAnswerColumn    = "answer"
	cardFileTypeColumn  = "file_type"
	cardFileIdColumn    = "file_id"
	cardCreatedAtColumn = "created_at"
	cardUpdatedAtColumn = "updated_at"
)

type CardRepository struct {
	dbal *dbal.DBAL
}

func NewCardRepository(dbal *dbal.DBAL) *CardRepository {
	return &CardRepository{
		dbal: dbal,
	}
}

func (r *CardRepository) Create(ctx context.Context, model *entity.Card) error {
	sql, args, err := r.dbal.SqlBuilder().
		Insert(cardTableName).
		Columns(
			cardIdColumn,
			cardUserIdColumn,
			cardQuestionColumn,
			cardAnswerColumn,
			cardFileTypeColumn,
			cardFileIdColumn,
			cardCreatedAtColumn,
			cardUpdatedAtColumn,
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
		ToSql()

	if err != nil {
		return fmt.Errorf("build sql: %w", err)
	}

	_, err = r.dbal.Exec(ctx, sql, args...)

	if err != nil {
		return fmt.Errorf("execute sql \"%s\": %w", sql, err)
	}

	return nil
}
