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

func (s *CardImpl) Create(ctx context.Context, request *card.CreateRequest) (*card.CreateResponse, error) {
	s.log.InfoCtx(
		ctx,
		"incoming request",
		zap.String("front_content", request.FrontContent),
		zap.String("back_content", request.BackContent),
	)

	cmd := command.CardCreateCmd{
		FrontContent: request.FrontContent,
		BackContent:  request.BackContent,
	}

	e, err := s.cardCreateHandler.Handle(ctx, cmd)

	if err != nil {
		return nil, s.handleError(ctx, err)
	}

	return &card.CreateResponse{
		Card: &card.Card{
			Id:           e.Id().String(),
			FrontContent: e.FrontContent(),
			BackContent:  e.BackContent(),
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
