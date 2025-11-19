package bot_handler

import (
	"context"
	"telegram-bot/pkg/flow"
)

const cardInputDescriptionStep = "card.input_description"

func (h *BotHandler) createCardInputDescriptionStep() (string, flow.FlowFunction) {
	return cardInputDescriptionStep, func(ctx context.Context, step, prevStep string, args []any) {
		//if prevStep != "" {
		//	return
		//}
		//
		//_, update := h.getArgs(args)
		//
		//h.user[update.Message.From.ID].Meta["card.description"] = update.Message.Text
		//
		//h.f.Goto(ctx, cardCreateRecapStep)
	}
}
