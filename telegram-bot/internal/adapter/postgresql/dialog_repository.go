package postgresql

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"platform/pkg/dbal"
	"telegram-bot/internal/domain/entity"
)

const (
	dialogTableName       = "dialog"
	dialogIdColumn        = "id"
	dialogStepColumn      = "step"
	dialogParamsColumn    = "params"
	dialogUserIdColumn    = "user_id"
	dialogCreatedAtColumn = "created_at"
	dialogUpdatedAtColumn = "updated_at"
)

type DialogRepository struct {
	dbal *dbal.DBAL
}

func NewDialogRepository(dbal *dbal.DBAL) *DialogRepository {
	return &DialogRepository{
		dbal: dbal,
	}
}

func (r *DialogRepository) FindByChatId(ctx context.Context, userId uuid.UUID) (*entity.Dialog, error) {
	sql, args, err := r.dbal.SqlBuilder().
		Select(
			dialogIdColumn,
			dialogStepColumn,
			dialogParamsColumn,
			dialogUserIdColumn,
			dialogCreatedAtColumn,
			dialogUpdatedAtColumn,
		).
		From(dialogTableName).
		Where(sq.Eq{dialogUserIdColumn: userId}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("build sql: %w", err)
	}

	m := &entity.Dialog{}

	err = r.dbal.ScanOne(ctx, m, sql, args...)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("scan one: %w", err)
	}

	return m, nil
}

func (r *DialogRepository) Create(ctx context.Context, model *entity.Dialog) error {
	sql, args, err := r.dbal.SqlBuilder().
		Insert(dialogTableName).
		Columns(
			dialogIdColumn,
			dialogStepColumn,
			dialogParamsColumn,
			dialogUserIdColumn,
			dialogCreatedAtColumn,
			dialogUpdatedAtColumn,
		).
		Values(
			model.Id,
			model.Step,
			model.Params,
			model.UserId,
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

func (r *DialogRepository) Update(ctx context.Context, model *entity.Dialog) error {
	sql, args, err := r.dbal.SqlBuilder().
		Update(dialogTableName).
		Set(dialogStepColumn, model.Step).
		Set(dialogParamsColumn, model.Params).
		Set(dialogUpdatedAtColumn, model.UpdatedAt).
		Where(sq.Eq{dialogIdColumn: model.Id}).
		ToSql()

	if err != nil {
		return fmt.Errorf("build sql: %w", err)
	}

	_, err = r.dbal.Exec(ctx, sql, args...)

	if err != nil {
		return fmt.Errorf("execute sql: \"%s\": %w", sql, err)
	}

	return nil
}
