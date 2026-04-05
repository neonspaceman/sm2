package postgresql

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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

func (r *UserRepository) FindByChatId(ctx context.Context, chatId int64) (*entity.User, error) {
	sql, args := selectUser.
		Where(sq.Eq{userChatIdColumn: chatId}).
		MustSql()

	rows, err := r.conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("query_builder: %w", err)
	}

	model, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[entity.User])

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("collect one: %w", err)
	}

	return model, nil
}

func (r *UserRepository) Create(ctx context.Context, model *entity.User) error {
	sql, args := sq.
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
		MustSql()

	fmt.Println(sql, args)

	_, err := r.conn.Exec(ctx, sql, args...)

	if err != nil {
		return fmt.Errorf("execute sql \"%s\": %w", sql, err)
	}

	return nil
}
