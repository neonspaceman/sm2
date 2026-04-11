package grpc

import (
	"card/internal/grpc/mappers"
	"card/internal/usecase/query"
	card_api "card/pkg/api/card"
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (s *CardImpl) GetByUserId(ctx context.Context, req *card_api.GetByUserIdRequest) (*card_api.GetByUserIdResponse, error) {
	s.log.InfoCtx(
		ctx,
		"Incoming get by user id request",
		zap.String("user_id", req.UserId),
	)

	cmd := query.GetCardsByUserIdQuery{
		UserId: req.UserId,
		Limit:  req.Limit + 1,
		After:  req.After,
	}

	cards, err := s.getCardsByUserIdQuery.Handle(ctx, cmd)

	if err != nil {
		return nil, s.handleError(ctx, err)
	}

	var encCursor uuid.UUID
	res := make([]*card_api.Card, 0, req.Limit)
	limit := min(uint64(len(cards)), req.Limit)

	for _, card := range cards[:limit] {
		res = append(res, mappers.ToCard(card))
		encCursor = card.Id
	}

	return &card_api.GetByUserIdResponse{
		Cards:     res,
		HasNext:   uint64(len(cards)) > req.Limit,
		EndCursor: encCursor.String(),
	}, nil
}
