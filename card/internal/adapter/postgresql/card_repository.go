package postgresql

import (
	"card/internal/consts"
	"card/internal/domain/entity"
	"card/pkg/dbal"
	"context"
	"fmt"
)

type CardRepository struct {
	dbal *dbal.DBAL
}

func NewCardRepository(dbal *dbal.DBAL) *CardRepository {
	return &CardRepository{
		dbal: dbal,
	}
}

func (r *CardRepository) Create(ctx context.Context, model entity.CardCard) error {
	sql, args, err := r.dbal.SqlBuilder().
		Insert(consts.CardTableName).
		Columns(
			consts.CardIdColumn,
			consts.CardFrontContentColumn,
			consts.CardBackContentColumn,
			consts.CardCreatedAtColumn,
			consts.CardUpdatedAtColumn,
		).
		Values(
			model.Id(),
			model.FrontContent(),
			model.BackContent(),
			model.CreatedAt(),
			model.UpdatedAt(),
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("NewCardRepository.Create: unable build sql: %w", err)
	}

	_, err = r.dbal.Exec(ctx, sql, args...)

	if err != nil {
		return fmt.Errorf("NewCardRepository.Create: unable to execute sql \"%s\": %w", sql, err)
	}

	return nil
}
