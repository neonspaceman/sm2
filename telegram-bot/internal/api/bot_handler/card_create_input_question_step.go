package bot_handler

import (
	"context"
	"github.com/go-telegram/bot"
	"go.uber.org/zap"
	"telegram-bot/internal/domain/types"
	"telegram-bot/internal/usercase/command"
	"telegram-bot/pkg/flow"
)

const cardCreateInputQuestionStep = "card_create.input_question"

func (h *BotHandler) createCardCreateInputQuestionStep() (string, flow.FlowFunction) {
	return cardCreateInputQuestionStep, func(ctx context.Context, step, prevStep string, args []any) {
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
				types.DialogParamCardQuestion: state.Update.Message.Text,
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

		h.f.Goto(ctx, cardCreateImageStep)
	}
}
