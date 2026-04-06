package repository

import (
	"card/internal/domain/entity"
	"context"
)

type CardStateRepositoryInterface interface {
	Create(ctx context.Context, model *entity.CardState) error
}
