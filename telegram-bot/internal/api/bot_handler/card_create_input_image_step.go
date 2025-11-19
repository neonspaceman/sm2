package bot_handler

import (
	"context"
	"github.com/go-telegram/bot"
	"go.uber.org/zap"
	"telegram-bot/internal/client/card"
	"telegram-bot/internal/domain/types"
	"telegram-bot/internal/usercase/command"
	"telegram-bot/pkg/flow"
)

const cardCreateInputImageStep = "card_create.input_image"

func (h *BotHandler) createCardCreateInputImageStep() (string, flow.FlowFunction) {
	return cardCreateInputImageStep, func(ctx context.Context, step, prevStep string, args []any) {
		if prevStep != "" {
			return
		}

		state := h.getState(args)

		fileType := card.FileTypeNone
		fileId := ""

		switch {
		case len(state.Update.Message.Photo) > 0:
			fileType = card.FileTypePhoto
			fileId = state.Update.Message.Photo[0].FileID
		case state.Update.Message.Document != nil:
			fileType = card.FileTypeDocument
			fileId = state.Update.Message.Document.FileID
		}

		err := h.dialogUpdateHandler.Handle(ctx, command.DialogUpdateCmd{
			Dialog: state.Dialog,
			Params: types.DialogParams{
				types.DialogParamCardFileId:   fileId,
				types.DialogParamCardFileType: string(fileType),
			},
		})

		if err != nil {
			h.async(func() {
				_, err := state.Bot.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: state.User.ChatId,
					Text:   "Something went wrong, try again",
				})

				if err != nil {
					h.log.Error("Unable to send message", zap.Error(err))
				}
			})

			h.log.Error("Unable to update dialog", zap.Error(err))

			return
		}

		h.f.Goto(ctx, cardCreateAnswerStep)
	}
}
