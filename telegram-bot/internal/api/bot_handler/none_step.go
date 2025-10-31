package bot_handler

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegram-bot/pkg/flow"
)

const noneStep = "none"

func (h *BotHandler) createNoneStep() (string, flow.FlowFunction) {
	return noneStep, func(ctx context.Context, step, prevStep string, args []any) {
		state := h.getState(args)

		message := &bot.SendMessageParams{
			ChatID: state.User.ChatId,
			Text:   "Welcome",
			ReplyMarkup: &models.ReplyKeyboardMarkup{
				Keyboard: [][]models.KeyboardButton{
					{
						{Text: "Some nice button"},
					},
				},
				ResizeKeyboard: true,
			},
		}

		switch state.Update.Message.Text {
		case "Some nice button":
			message = &bot.SendMessageParams{
				ChatID: state.User.ChatId,
				Text:   "Welcome",
				ReplyMarkup: &models.ReplyKeyboardMarkup{
					Keyboard: [][]models.KeyboardButton{
						{
							{Text: "Some nice button"},
						},
					},
					ResizeKeyboard: true,
				},
			}
		}

		h.async(func() {
			_, _ = state.Bot.SendMessage(ctx, message)
		})
	}
}
