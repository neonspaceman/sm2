package grpc

import (
	"card/internal/usecase/command"
	"card/internal/usecase/query"
	"card/pkg/api/card"
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
