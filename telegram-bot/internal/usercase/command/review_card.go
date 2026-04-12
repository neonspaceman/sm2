package command

import (
	"context"
	"github.com/google/uuid"
	card_client "telegram-bot/internal/client/card"
	"telegram-bot/internal/domain/user"
)

type ReviewCardCmd struct {
	User   *user.User
	CardId uuid.UUID
	Rating string
}

type ReviewCardHandler struct {
	cardClient card_client.CardClientInterface
}

func NewReviewCardHandler(
	cardClient card_client.CardClientInterface,
) *ReviewCardHandler {
	return &ReviewCardHandler{
		cardClient: cardClient,
	}
}

func (h *ReviewCardHandler) Handle(ctx context.Context, cmd ReviewCardCmd) (*card_client.ReviewLog, error) {
	reviewLog, err := h.cardClient.ReviewCard(ctx, cmd.User.Id, cmd.CardId, cmd.Rating)
	if err != nil {
		return nil, err
	}

	return reviewLog, nil
}
