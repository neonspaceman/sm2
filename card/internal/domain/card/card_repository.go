package card

import (
	"context"
)

type CardRepositoryInterface interface {
	Create(ctx context.Context, model *Card) error
}
