package bot_handler

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
	"telegram-bot/pkg/flow"
)

const noneStep = "none"

const keyboardNewCard = "New card"

func (h *BotHandler) createNoneStep() (string, flow.FlowFunction) {
	return noneStep, func(ctx context.Context, step, prevStep string, args []any) {
		state := h.getState(args)

		switch state.Update.Message.Text {
		case keyboardNewCard:
			h.f.Goto(ctx, cardCreateStep)
		default:
			message := &bot.SendMessageParams{
				ChatID: state.User.ChatId,
				Text:   "Choose the option",
				ReplyMarkup: &models.ReplyKeyboardMarkup{
					Keyboard: [][]models.KeyboardButton{
						{
							{Text: keyboardNewCard},
						},
					},
					ResizeKeyboard: true,
				},
			}

			h.async(func() {
				_, err := state.Bot.SendMessage(ctx, message)

				if err != nil {
					h.log.Error("Unable to send message", zap.Error(err))
				}
			})
		}
	}
}
