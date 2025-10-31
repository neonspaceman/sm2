package repository

import (
	"context"
	"telegram-bot/internal/domain/entity"
)

type UserRepositoryInterface interface {
	FindByChatId(ctx context.Context, chatId int64) (*entity.User, error)
	Create(ctx context.Context, model *entity.User) error
}
