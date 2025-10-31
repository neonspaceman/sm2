package bot_handler

import (
	"context"
	"telegram-bot/pkg/flow"
)

func (h *BotHandler) createBefore() flow.FlowFunction {
	return func(ctx context.Context, step, prevStep string, param []any) {

	}
}
