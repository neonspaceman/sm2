package bot_handler

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
	"telegram-bot/pkg/flow"
)

const cardCreateImageStep = "card_create.image"

func (h *BotHandler) createCardCreateImageStep() (string, flow.FlowFunction) {
	return cardCreateImageStep, func(ctx context.Context, step, prevStep string, args []any) {
		state := h.getState(args)

		h.async(func() {
			_, err := state.Bot.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: state.User.ChatId,
				Text:   "Enter the image via @pic or @gif",
				ReplyMarkup: &models.ReplyKeyboardMarkup{
					Keyboard: [][]models.KeyboardButton{
						{
							{Text: "Skip the image step"},
						},
						{
							{Text: "Cancel"},
						},
					},
					ResizeKeyboard: true,
				},
			})

			if err != nil {
				h.log.Error("Unable to send message", zap.Error(err))
			}
		})

		h.f.Goto(ctx, cardCreateInputImageStep)
	}
}
