package bot_handler

import (
	"context"
	"go.uber.org/zap"
	"telegram-bot/internal/usercase/command"
	"telegram-bot/pkg/flow"
)

func (h *BotHandler) createAfter() flow.FlowFunction {
	return func(ctx context.Context, step, prevStep string, args []any) {
		state := h.getState(args)

		err := h.dialogUpdateHandler.Handle(ctx, command.DialogUpdateCmd{
			Dialog: state.Dialog,
			Step:   step,
		})

		if err != nil {
			h.log.Error("Unable to update dialog", zap.Error(err), zap.String("id", state.Dialog.Id.String()))
		}
	}
}
