package card

import (
	"context"
	"github.com/google/uuid"
)

type CardClientInterface interface {
	GetCards(ctx context.Context, userId uuid.UUID, limit uint64, cursor string) ([]*Card, bool, string, error)
	Create(ctx context.Context, userId uuid.UUID, question, answer string) (*Card, error)
}
