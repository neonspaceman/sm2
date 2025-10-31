package bot_handler

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
	"platform/pkg/logger"
	"telegram-bot/internal/domain/entity"
	"telegram-bot/internal/usercase/command"
	"telegram-bot/pkg/flow"
)

type stateType struct {
	Bot    *bot.Bot
	Dialog *entity.Dialog
	Update *models.Update
	User   *entity.User
}

type BotHandler struct {
	userFirstOrCreateHandler   *command.UserFirstOrCreateHandler
	dialogFirstOrCreateHandler *command.DialogFirstOrCreateHandler
	dialogUpdateHandler        *command.DialogUpdateHandler
	log                        *logger.Logger

	f     *flow.Flow
	ch    chan func()
	close chan struct{}
}

func NewBotHandler(
	userFirstOrCreateHandler *command.UserFirstOrCreateHandler,
	dialogFirstOrCreateHandler *command.DialogFirstOrCreateHandler,
	dialogUpdateHandler *command.DialogUpdateHandler,
	log *logger.Logger,
) *BotHandler {
	b := &BotHandler{
		userFirstOrCreateHandler:   userFirstOrCreateHandler,
		dialogFirstOrCreateHandler: dialogFirstOrCreateHandler,
		dialogUpdateHandler:        dialogUpdateHandler,
		log:                        log,
	}

	f := flow.NewBuilder().
		Before(b.createBefore()).
		Step(b.createNoneStep()).
		After(b.createAfter()).
		UnhandledState(func(ctx context.Context, step, prevStep string, args []any) {
			fmt.Printf("Unhandled step: %s\n", step)
		}).
		Flow()

	b.f = f
	b.ch = make(chan func(), 100)

	// Async worker
	go func() {
		defer close(b.ch)

		for {
			select {
			case callback, ok := <-b.ch:
				if !ok {
					return
				}

				callback()
			case <-b.close:
				return
			}
		}
	}()

	return b
}

func (h *BotHandler) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	user, err := h.userFirstOrCreateHandler.Handle(ctx, command.UserFirstOrCreateCmd{
		ChatId:    update.Message.From.ID,
		FirstName: update.Message.From.FirstName,
	})

	if err != nil {
		h.log.Error("Unable to get chat", zap.Error(err), zap.Int64("chat_id", update.Message.From.ID))

		return
	}

	dialog, err := h.dialogFirstOrCreateHandler.Handle(ctx, command.DialogFirstOrCreateCmd{
		ChatId: user.Id,
	})

	if err != nil {
		h.log.Error("Unable to get dialog", zap.Error(err), zap.String("user_id", user.Id.String()))
	}

	h.f.Active(ctx, dialog.Step, &stateType{
		Bot:    b,
		Dialog: dialog,
		Update: update,
		User:   user,
	})
}

func (h *BotHandler) Close() {
	h.close <- struct{}{}
}

func (h *BotHandler) async(callback func()) {
	h.ch <- callback
}

func (h *BotHandler) getState(args []any) *stateType {
	return args[0].(*stateType)
}
