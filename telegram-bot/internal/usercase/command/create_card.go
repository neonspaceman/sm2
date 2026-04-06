package command

import (
	"context"
	card_client "telegram-bot/internal/client/card"
	"telegram-bot/internal/domain/entity"
)

type CreateCardCmd struct {
	User     *entity.User
	Answer   string
	Question string
}

type CreateCardHandler struct {
	cardClient card_client.CardClientInterface
}

func NewCreateCardHandler(
	cardClient card_client.CardClientInterface,
) *CreateCardHandler {
	return &CreateCardHandler{
		cardClient: cardClient,
	}
}

func (h *CreateCardHandler) Handle(ctx context.Context, cmd CreateCardCmd) (*card_client.Card, error) {
	createdCard, err := h.cardClient.Create(ctx, cmd.User.Id, cmd.Question, cmd.Answer)

	if err != nil {
		return nil, err
	}

	return createdCard, nil
}
