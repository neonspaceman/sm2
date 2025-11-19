package bot_handler

import (
	"context"
	"github.com/go-telegram/bot"
	"go.uber.org/zap"
	"telegram-bot/internal/domain/types"
	"telegram-bot/internal/usercase/command"
	"telegram-bot/pkg/flow"
)

const cardCreateInputAnswerStep = "card_create.input_answer"

func (h *BotHandler) createCardCreateInputAnswerStep() (string, flow.FlowFunction) {
	return cardCreateInputAnswerStep, func(ctx context.Context, step, prevStep string, args []any) {
		if prevStep != "" {
			return
		}

		state := h.getState(args)

		if state.Update.Message.Text == "" {
			h.async(func() {
				_, err := state.Bot.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: state.User.ChatId,
					Text:   "Enter the text",
				})

				if err != nil {
					h.log.Error("Unable to send message", zap.Error(err))
				}
			})

			return
		}

		err := h.dialogUpdateHandler.Handle(ctx, command.DialogUpdateCmd{
			Dialog: state.Dialog,
			Params: types.DialogParams{
				types.DialogParamCardAnswer: state.Update.Message.Text,
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

		h.f.Goto(ctx, cardCreateRecapStep)
	}
}
