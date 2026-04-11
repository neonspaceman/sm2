package review

import (
	"context"
)

type ReviewLogRepositoryInterface interface {
	Create(ctx context.Context, model *ReviewLog) error
}
