package grpc

import (
	"card/internal/grpc/mappers"
	"card/internal/usecase/command"
	cardPkg "card/pkg/api/card"
	"context"
	"go.uber.org/zap"
)

func (c *CardImpl) Create(ctx context.Context, req *cardPkg.CreateRequest) (*cardPkg.CreateResponse, error) {
	c.log.InfoCtx(
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

	card, err := c.cardCreateHandler.Handle(ctx, cmd)

	if err != nil {
		return nil, c.handleError(ctx, err)
	}

	return &cardPkg.CreateResponse{Card: mappers.ToCard(card)}, nil
}
