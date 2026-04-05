package command

import (
	"context"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"telegram-bot/internal/domain/entity"
	"telegram-bot/internal/domain/repository"
)

type UserFirstOrCreateHandler struct {
	repository repository.UserRepositoryInterface
}

func NewUserFirstOrCreateHandler(
	repository repository.UserRepositoryInterface,
) *UserFirstOrCreateHandler {
	return &UserFirstOrCreateHandler{
		repository: repository,
	}
}

func (h *UserFirstOrCreateHandler) Handle(ctx context.Context, data *initdata.InitData) (*entity.User, error) {
	u, err := h.repository.FindByChatId(ctx, data.Chat.ID)

	if err != nil {
		return nil, err
	}

	if u != nil {
		return u, nil
	}

	u = entity.NewUser(data.Chat.ID, data.User.FirstName)

	err = h.repository.Create(ctx, u)

	if err != nil {
		return nil, err
	}

	return u, nil
}
