package card

import "context"

type CardClientInterface interface {
	Create(ctx context.Context, req CreateRequestDto) error
}
