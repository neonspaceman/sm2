package grpc

import (
	"card/internal/usecase/command"
	"card/pkg/api/card"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"platform/pkg/logger"
)

type CardImplProps struct {
	Log               *logger.Logger
	CardCreateHandler *command.CardCreateHandler
}

type CardImpl struct {
	card.UnimplementedCardServiceServer
	log               *logger.Logger
	cardCreateHandler *command.CardCreateHandler
}

func NewCardImpl(props CardImplProps) *CardImpl {
	return &CardImpl{
		log:               props.Log,
		cardCreateHandler: props.CardCreateHandler,
	}
}

func (s *CardImpl) Create(ctx context.Context, req *card.CreateRequest) (*card.CreateResponse, error) {
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

	e, err := s.cardCreateHandler.Handle(ctx, cmd)

	if err != nil {
		return nil, s.handleError(ctx, err)
	}

	return &card.CreateResponse{
		Card: &card.Card{
			Id:       e.Id.String(),
			Question: e.Question,
			Answer:   e.Answer,
			FileType: card.FileType(card.FileType_value[string(e.FileType)]),
			FileId:   e.FileId,
		},
	}, nil
}

func (s *CardImpl) handleError(ctx context.Context, err error) error {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		grpcStatus := status.New(codes.InvalidArgument, validationErrors.Error())
		grpcStatus, err = grpcStatus.WithDetails(&errdetails.ErrorInfo{
			Reason: "validation_error",
		})

		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		return grpcStatus.Err()
	}

	s.log.ErrorCtx(ctx, "internal error", zap.String("error", err.Error()))

	return status.Error(codes.Internal, err.Error())
}
