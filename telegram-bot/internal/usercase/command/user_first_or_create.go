package command

import (
	"context"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"telegram-bot/internal/domain/user"
)

type FirstOrCreateUserHandler struct {
	repository user.UserRepositoryInterface
}

func NewFirstOrCreateUserHandler(
	repository user.UserRepositoryInterface,
) *FirstOrCreateUserHandler {
	return &FirstOrCreateUserHandler{
		repository: repository,
	}
}

func (h *FirstOrCreateUserHandler) Handle(ctx context.Context, data initdata.InitData) (*user.User, error) {
	u, err := h.repository.FindByChatId(ctx, data.Chat.ID)

	if err != nil {
		return nil, err
	}

	if u != nil {
		return u, nil
	}

	u = user.NewUser(data.Chat.ID, data.User.FirstName)

	err = h.repository.Create(ctx, u)

	if err != nil {
		return nil, err
	}

	return u, nil
}
