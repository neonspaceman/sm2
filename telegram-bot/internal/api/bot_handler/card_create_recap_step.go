package bot_handler

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
	"telegram-bot/internal/client/card"
	"telegram-bot/internal/domain/types"
	"telegram-bot/pkg/flow"
)

const cardCreateRecapStep = "card_create.recap"

func (h *BotHandler) createCardCreateRecapStep() (string, flow.FlowFunction) {
	return cardCreateRecapStep, func(ctx context.Context, step, prevStep string, args []any) {
		state := h.getState(args)

		question := state.Dialog.Params[types.DialogParamCardQuestion]
		fileId := state.Dialog.Params[types.DialogParamCardFileId]
		fileType := card.FileType(state.Dialog.Params[types.DialogParamCardFileType])
		answer := state.Dialog.Params[types.DialogParamCardAnswer]

		switch fileType {
		case card.FileTypePhoto:
			h.async(func() {
				_, err := state.Bot.SendPhoto(ctx, &bot.SendPhotoParams{
					ChatID:  state.User.ChatId,
					Photo:   &models.InputFileString{Data: fileId},
					Caption: fmt.Sprintf("Q: %s\nA: %s", question, answer),
					ReplyMarkup: &models.ReplyKeyboardMarkup{
						Keyboard: [][]models.KeyboardButton{
							{
								{Text: "Save"},
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
		case card.FileTypeDocument:
			h.async(func() {
				_, err := state.Bot.SendDocument(ctx, &bot.SendDocumentParams{
					ChatID:   state.User.ChatId,
					Document: &models.InputFileString{Data: fileId},
					Caption:  fmt.Sprintf("Q: %s\nA: %s", question, answer),
					ReplyMarkup: &models.ReplyKeyboardMarkup{
						Keyboard: [][]models.KeyboardButton{
							{
								{Text: "Save"},
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
		default:
			h.async(func() {
				_, err := state.Bot.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: state.User.ChatId,
					Text:   fmt.Sprintf("Q: %s\nA: %s", question, answer),
					ReplyMarkup: &models.ReplyKeyboardMarkup{
						Keyboard: [][]models.KeyboardButton{
							{
								{Text: "Save"},
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
		}

		h.f.Goto(ctx, cardCreateDoneStep)
	}
}
