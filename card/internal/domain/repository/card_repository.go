package repository

import (
	"card/internal/domain/entity"
	"context"
)

type CardRepositoryInterface interface {
	Create(ctx context.Context, model entity.CardCard) error
}
