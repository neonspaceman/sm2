package bot_handler

import (
	"context"
	"telegram-bot/internal/client/card"
	"telegram-bot/internal/domain/types"
	"telegram-bot/pkg/flow"
)

const cardCreateDoneStep = "card_create.done"

func (h *BotHandler) createCardCreateDoneStep() (string, flow.FlowFunction) {
	return cardCreateDoneStep, func(ctx context.Context, step, prevStep string, args []any) {
		if prevStep != "" {
			return
		}

		state := h.getState(args)

		if state.Update.Message.Text == "Save" {
			question := state.Dialog.Params[types.DialogParamCardQuestion]
			answer := state.Dialog.Params[types.DialogParamCardAnswer]
			fileType := card.FileType(state.Dialog.Params[types.DialogParamCardFileType])
			fileId := state.Dialog.Params[types.DialogParamCardFileId]

			err := h.cardClient.Create(ctx, card.CreateRequestDto{
				UserId:   state.User.Id.String(),
				Question: question,
				Answer:   answer,
				FileType: fileType,
				FileId:   fileId,
			})

			// TODO: Error handler
			if err != nil {
				panic(err)
			}

			//h.log.Info("Save card", zap.String("file_id", fileId), zap.String("question", question))
			//
			//h.async(func() {
			//	_, err := state.Bot.SendMessage(ctx, &bot.SendMessageParams{
			//		ChatID: state.User.ChatId,
			//		Text:   "The new card has been added",
			//	})
			//
			//	if err != nil {
			//		h.log.Error("Unable to send message", zap.Error(err))
			//	}
			//})
		}

		h.f.Goto(ctx, noneStep)
	}
}
