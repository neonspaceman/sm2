package grpc

import (
	"card/internal/grpc/mappers"
	"card/internal/usecase/command"
	card_api "card/pkg/api/card"
	"context"
	"go.uber.org/zap"
)

func (s *CardImpl) Create(ctx context.Context, req *card_api.CreateRequest) (*card_api.CreateResponse, error) {
	s.log.InfoCtx(
		ctx,
		"Incoming create request",
		zap.String("question", req.Question),
		zap.String("answer", req.Answer),
		zap.String("file_type", req.FileType.String()),
		zap.String("file_id", req.FileId),
	)

	cmd := command.CreateCardCmd{
		UserId:   req.UserId,
		Question: req.Question,
		Answer:   req.Answer,
		FileType: req.FileType.String(),
		FileId:   req.FileId,
	}

	card, err := s.cardCreateHandler.Handle(ctx, cmd)

	if err != nil {
		return nil, s.handleError(ctx, err)
	}

	return &card_api.CreateResponse{Card: mappers.ToCard(card)}, nil
}
