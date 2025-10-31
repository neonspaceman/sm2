package command

import (
	"context"
	"telegram-bot/internal/domain/entity"
	"telegram-bot/internal/domain/repository"
)

type DialogUpdateCmd struct {
	Dialog *entity.Dialog
	Step   string
	Params map[string]string
}

type DialogUpdateHandler struct {
	repository repository.DialogRepositoryInterface
}

func NewDialogUpdateHandler(
	repository repository.DialogRepositoryInterface,
) *DialogUpdateHandler {
	return &DialogUpdateHandler{
		repository: repository,
	}
}

func (h *DialogUpdateHandler) Handle(ctx context.Context, cmd DialogUpdateCmd) error {
	if cmd.Step != "" {
		cmd.Dialog.SetStep(cmd.Step)
	}

	if len(cmd.Params) > 0 {
		for k, v := range cmd.Params {
			cmd.Dialog.SetParam(k, v)
		}
	}

	err := h.repository.Update(ctx, cmd.Dialog)

	if err != nil {
		return err
	}

	return nil
}
