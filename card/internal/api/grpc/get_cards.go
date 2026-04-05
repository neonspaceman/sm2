package grpc

import (
	"card/internal/grpc/mappers"
	"card/internal/usecase/query"
	cardPkg "card/pkg/api/card"
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (c *CardImpl) GetByUserId(ctx context.Context, req *cardPkg.GetByUserIdRequest) (*cardPkg.GetByUserIdResponse, error) {
	c.log.InfoCtx(
		ctx,
		"Incoming get by user id request",
		zap.String("user_id", req.UserId),
	)

	cmd := query.GetCardsByUserIdQuery{
		UserId: req.UserId,
		Limit:  req.Limit + 1,
		After:  req.After,
	}

	cards, err := c.getCardsByUserIdQuery.Handle(ctx, cmd)

	if err != nil {
		return nil, c.handleError(ctx, err)
	}

	var encCursor uuid.UUID
	res := make([]*cardPkg.Card, 0, req.Limit)
	limit := min(uint64(len(cards)), req.Limit)

	for _, card := range cards[:limit] {
		res = append(res, mappers.ToCard(card))
		encCursor = card.Id
	}

	return &cardPkg.GetByUserIdResponse{
		Cards:     res,
		HasNext:   uint64(len(cards)) > req.Limit,
		EndCursor: encCursor.String(),
	}, nil
}
