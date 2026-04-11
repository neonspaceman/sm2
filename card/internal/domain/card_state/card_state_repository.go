package card_state

import (
	"context"
	"github.com/google/uuid"
)

type CardStateRepositoryInterface interface {
	Create(ctx context.Context, model *CardState) error
	GetById(ctx context.Context, id uuid.UUID) (*CardState, error)
}
