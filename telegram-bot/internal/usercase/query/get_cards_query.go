package query

import (
	"context"
	"telegram-bot/internal/adapter/grpc"
	"telegram-bot/internal/client/card"
	"telegram-bot/internal/domain/entity"
)

type GetCardsByUserIdQuery struct {
	User  *entity.User
	Limit uint64
	After string
}

type GetCardByUserIdHandler struct {
	cardClient *grpc.CardClient
}

func NewGetCardByUserIdHandler(
	cardClient *grpc.CardClient,
) *GetCardByUserIdHandler {
	return &GetCardByUserIdHandler{
		cardClient: cardClient,
	}
}

func (h *GetCardByUserIdHandler) Handle(ctx context.Context, cmd GetCardsByUserIdQuery) ([]*card.Card, bool, string, error) {
	return h.cardClient.GetCards(ctx, cmd.User.Id, cmd.Limit, cmd.After)
}
