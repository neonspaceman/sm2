package command

import (
	"context"
	"github.com/google/uuid"
	"telegram-bot/internal/domain/entity"
	"telegram-bot/internal/domain/repository"
	"telegram-bot/internal/domain/types"
)

const defaultStep = "none"

type DialogFirstOrCreateCmd struct {
	ChatId uuid.UUID
}

type DialogFirstOrCreateHandler struct {
	repository repository.DialogRepositoryInterface
}

func NewDialogFirstOrCreateHandler(
	repository repository.DialogRepositoryInterface,
) *DialogFirstOrCreateHandler {
	return &DialogFirstOrCreateHandler{
		repository: repository,
	}
}

func (h *DialogFirstOrCreateHandler) Handle(ctx context.Context, cmd DialogFirstOrCreateCmd) (*entity.Dialog, error) {
	d, err := h.repository.FindByChatId(ctx, cmd.ChatId)

	if err != nil {
		return nil, err
	}

	if d != nil {
		return d, nil
	}

	d = entity.NewDialog(defaultStep, types.DialogParams{}, cmd.ChatId)

	err = h.repository.Create(ctx, d)

	if err != nil {
		return nil, err
	}

	return d, nil
}
