package query

import (
	"context"
	card_client "telegram-bot/internal/client/card"
	"telegram-bot/internal/domain/user"
)

type GetCardsByUserIdQuery struct {
	User  *user.User
	Limit uint64
	After string
}

type GetCardByUserIdHandler struct {
	cardClient card_client.CardClientInterface
}

func NewGetCardByUserIdHandler(
	cardClient card_client.CardClientInterface,
) *GetCardByUserIdHandler {
	return &GetCardByUserIdHandler{
		cardClient: cardClient,
	}
}

func (h *GetCardByUserIdHandler) Handle(ctx context.Context, cmd GetCardsByUserIdQuery) ([]*card_client.Card, bool, string, error) {
	return h.cardClient.GetCards(ctx, cmd.User.Id, cmd.Limit, cmd.After)
}
