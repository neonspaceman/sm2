package card

import (
	"context"
	"github.com/google/uuid"
)

type CardRepositoryInterface interface {
	Create(ctx context.Context, model *Card) error
	GetById(ctx context.Context, id uuid.UUID) (*Card, error)
}
