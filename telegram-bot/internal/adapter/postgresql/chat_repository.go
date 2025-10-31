package postgresql

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"platform/pkg/dbal"
	"telegram-bot/internal/domain/entity"
)

const (
	userTableName       = "\"user\""
	userIdColumn        = "id"
	userChatIdColumn    = "chat_id"
	userFirstNameColumn = "first_name"
	userCreatedAtColumn = "created_at"
)

type UserRepository struct {
	dbal *dbal.DBAL
}

func NewUserRepository(dbal *dbal.DBAL) *UserRepository {
	return &UserRepository{
		dbal: dbal,
	}
}

func (r *UserRepository) FindByChatId(ctx context.Context, chatId int64) (*entity.User, error) {
	sql, args, err := r.dbal.SqlBuilder().
		Select(
			userIdColumn,
			userChatIdColumn,
			userFirstNameColumn,
			userCreatedAtColumn,
		).
		From(userTableName).
		Where(sq.Eq{userChatIdColumn: chatId}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("build sql: %w", err)
	}

	m := &entity.User{}

	err = r.dbal.ScanOne(ctx, m, sql, args...)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("scan one: %w", err)
	}

	return m, nil
}

func (r *UserRepository) Create(ctx context.Context, model *entity.User) error {
	sql, args, err := r.dbal.SqlBuilder().
		Insert(userTableName).
		Columns(
			userIdColumn,
			userChatIdColumn,
			userFirstNameColumn,
			userCreatedAtColumn,
		).
		Values(
			model.Id,
			model.ChatId,
			model.FirstName,
			model.CreatedAt,
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
