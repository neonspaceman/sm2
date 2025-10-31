package command

import (
	"context"
	"github.com/go-playground/validator/v10"
	"telegram-bot/internal/domain/entity"
	"telegram-bot/internal/domain/repository"
)

type UserFirstOrCreateCmd struct {
	ChatId    int64  `validate:"required"`
	FirstName string `validate:"required"`
}

type UserFirstOrCreateHandler struct {
	repository repository.UserRepositoryInterface
	validate   *validator.Validate
}

func NewUserFirstOrCreateHandler(
	repository repository.UserRepositoryInterface,
	validate *validator.Validate,
) *UserFirstOrCreateHandler {
	return &UserFirstOrCreateHandler{
		repository: repository,
		validate:   validate,
	}
}

func (h *UserFirstOrCreateHandler) Handle(ctx context.Context, cmd UserFirstOrCreateCmd) (*entity.User, error) {
	err := h.validate.Struct(cmd)

	if err != nil {
		return nil, err
	}

	u, err := h.repository.FindByChatId(ctx, cmd.ChatId)

	if err != nil {
		return nil, err
	}

	if u != nil {
		return u, nil
	}

	u = entity.NewUser(cmd.ChatId, cmd.FirstName)

	err = h.repository.Create(ctx, u)

	if err != nil {
		return nil, err
	}

	return u, nil
}
