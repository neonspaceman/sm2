package postgresql

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	user_domain "telegram-bot/internal/domain/user"
)

const (
	userTableName       = "\"user\""
	userIdColumn        = "id"
	userChatIdColumn    = "chat_id"
	userFirstNameColumn = "first_name"
	userCreatedAtColumn = "created_at"
)

type UserRepository struct {
	conn *pgxpool.Pool
}

func NewUserRepository(conn *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

var selectUser = sq.
	Select(
		userIdColumn,
		userChatIdColumn,
		userFirstNameColumn,
		userCreatedAtColumn,
	).
	From(userTableName).
	PlaceholderFormat(sq.Dollar)

func (r *UserRepository) FindByChatId(ctx context.Context, chatId int64) (*user_domain.User, error) {
	sql, args, err := selectUser.
		Where(sq.Eq{userChatIdColumn: chatId}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("build get query: %w", err)
	}

	rows, err := r.conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("query_builder: %w", err)
	}

	model, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[user_domain.User])

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, user_domain.ErrUserNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("collect one: %w", err)
	}

	return model, nil
}

func (r *UserRepository) Create(ctx context.Context, model *user_domain.User) error {
	sql, args, err := sq.
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
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("build create query: %w", err)
	}

	_, err = r.conn.Exec(ctx, sql, args...)

	if err != nil {
		return fmt.Errorf("user repository: fail to insert user: %w", err)
	}

	return nil
}
