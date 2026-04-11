package grpc

import (
	"card/internal/usecase/command"
	"card/internal/usecase/query"
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
	Log                   *logger.Logger
	CreateCardHandler     *command.CardCreateHandler
	GetCardsByUserIdQuery *query.GetCardByUserIdHandler
	ReviewCardHandler     *command.ReviewCardHandler
}

type CardImpl struct {
	card.UnimplementedCardServiceServer
	log                   *logger.Logger
	cardCreateHandler     *command.CardCreateHandler
	getCardsByUserIdQuery *query.GetCardByUserIdHandler
	reviewCardHandler     *command.ReviewCardHandler
}

func NewCardImpl(props CardImplProps) *CardImpl {
	return &CardImpl{
		log:                   props.Log,
		cardCreateHandler:     props.CreateCardHandler,
		getCardsByUserIdQuery: props.GetCardsByUserIdQuery,
		reviewCardHandler:     props.ReviewCardHandler,
	}
}

// TODO: change to interceptor?
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
