package user

import (
	"context"
)

type UserRepositoryInterface interface {
	FindByChatId(ctx context.Context, chatId int64) (*User, error)
	Create(ctx context.Context, model *User) error
}
