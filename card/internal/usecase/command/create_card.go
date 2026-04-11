package command

import (
	domain_card "card/internal/domain/card"
	domain_card_state "card/internal/domain/card_state"
	"context"
	"fmt"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateCardCmd struct {
	UserId   string `validate:"required,uuid"`
	Question string `validate:"required"`
	Answer   string `validate:"required"`
	FileType string `validate:"required,oneof=NONE PHOTO DOCUMENT"`
	FileId   string
}

type CardCreateHandler struct {
	cardRepository      domain_card.CardRepositoryInterface
	cardStateRepository domain_card_state.CardStateRepositoryInterface
	trManager           *manager.Manager
	validate            *validator.Validate
}

func NewCreateCardHandler(
	cardRepository domain_card.CardRepositoryInterface,
	cardStateRepository domain_card_state.CardStateRepositoryInterface,
	trManager *manager.Manager,
	validate *validator.Validate,
) *CardCreateHandler {
	return &CardCreateHandler{
		cardRepository:      cardRepository,
		cardStateRepository: cardStateRepository,
		trManager:           trManager,
		validate:            validate,
	}
}

func (h *CardCreateHandler) Handle(ctx context.Context, cmd CreateCardCmd) (*domain_card.Card, error) {
	err := h.validate.Struct(&cmd)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(cmd.UserId)
	if err != nil {
		return nil, err
	}

	card, err := domain_card.NewCard(
		userId,
		cmd.Answer,
		cmd.Question,
		domain_card.FileType(cmd.FileType),
		cmd.FileId,
	)
	if err != nil {
		return nil, err
	}

	cardState := domain_card_state.NewCardState(card.Id)

	err = h.trManager.Do(ctx, func(ctx context.Context) error {
		err = h.cardRepository.Create(ctx, card)
		if err != nil {
			return err
		}

		err = h.cardStateRepository.Create(ctx, cardState)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("create new card: %w", err)
	}

	return card, nil
}
