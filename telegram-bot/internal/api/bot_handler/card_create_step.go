package bot_handler

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegram-bot/pkg/flow"
)

const cardCreateStep = "card_create"

func (h *BotHandler) createCardCreateStep() (string, flow.FlowFunction) {
	return cardCreateStep, func(ctx context.Context, step, prevStep string, args []any) {
		state := h.getState(args)

		h.async(func() {
			_, _ = state.Bot.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: state.User.ChatId,
				Text:   "Type in the question",
				ReplyMarkup: &models.ReplyKeyboardMarkup{
					Keyboard: [][]models.KeyboardButton{
						{
							{Text: "Cancel"},
						},
					},
					ResizeKeyboard: true,
				},
			})
		})

		h.f.Goto(ctx, cardCreateInputQuestionStep)
	}
}
