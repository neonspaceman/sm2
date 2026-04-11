package grpc

import (
	"card/internal/grpc/mappers"
	"card/internal/usecase/command"
	card_api "card/pkg/api/card"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (s *CardImpl) Review(ctx context.Context, req *card_api.ReviewCardRequest) (*card_api.ReviewCardResponse, error) {
	s.log.InfoCtx(
		ctx,
		"Incoming create request",
		zap.String("user_id", req.UserId),
		zap.String("card_id", req.CardId),
		zap.String("rating", req.Rating.String()),
		zap.String("review_at", req.ReviewAt.String()),
	)

	cardId, err := uuid.Parse(req.CardId)
	if err != nil {
		return nil, fmt.Errorf("parse card id: %w", err)
	}

	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("parse user id: %w", err)
	}

	rating, err := mappers.ToRating(req.Rating)
	if err != nil {
		return nil, fmt.Errorf("parse rating: %w", err)
	}

	reviewLog, err := s.reviewCardHandler.Handle(ctx, command.ReviewCardCmd{
		UserId:     userId,
		CardId:     cardId,
		ReviewedAt: req.ReviewAt.AsTime(),
		Rating:     rating,
	})
	if err != nil {
		return nil, fmt.Errorf("review: %w", err)
	}

	ratingResponse, err := mappers.FromRating(reviewLog.Rating)
	if err != nil {
		return nil, fmt.Errorf("parse rating: %w", err)
	}

	return &card_api.ReviewCardResponse{
		ReviewLog: &card_api.ReviewLog{
			Id:     reviewLog.Id.String(),
			CardId: reviewLog.CardId.String(),
			Rating: ratingResponse,
		},
	}, nil
}
