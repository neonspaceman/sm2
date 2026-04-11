package card_state

import (
	"context"
	"github.com/google/uuid"
)

type CardStateRepositoryInterface interface {
	GetById(ctx context.Context, id uuid.UUID) (*CardState, error)
	Create(ctx context.Context, model *CardState) error
	Save(ctx context.Context, model *CardState) error
}
