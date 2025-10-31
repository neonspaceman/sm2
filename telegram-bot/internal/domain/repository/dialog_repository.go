package repository

import (
	"context"
	"github.com/google/uuid"
	"telegram-bot/internal/domain/entity"
)

type DialogRepositoryInterface interface {
	FindByChatId(ctx context.Context, userId uuid.UUID) (*entity.Dialog, error)
	Create(ctx context.Context, model *entity.Dialog) error
	Update(ctx context.Context, model *entity.Dialog) error
}
