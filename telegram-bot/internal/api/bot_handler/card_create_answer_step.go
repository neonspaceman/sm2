package bot_handler

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegram-bot/pkg/flow"
)

const cardCreateAnswerStep = "card_create.answer"

func (h *BotHandler) createCardCreateAnswerStep() (string, flow.FlowFunction) {
	return cardCreateAnswerStep, func(ctx context.Context, step, prevStep string, args []any) {
		state := h.getState(args)

		h.async(func() {
			_, _ = state.Bot.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: state.User.ChatId,
				Text:   "Type in the answer",
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

		h.f.Goto(ctx, cardCreateInputAnswerStep)
	}
}
