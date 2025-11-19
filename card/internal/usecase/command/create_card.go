package command

import (
	entityPkg "card/internal/domain/entity"
	"card/internal/domain/repository"
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateCardCmd struct {
	UserId   string `validate:"required,uuid"`
	Question string `validate:"required"`
	Answer   string `validate:"required"`
	FileType string `validate:"required,oneof=NONE PHOTO DOCUMENT"`
	FileId   string
}

type CardCreateHandler struct {
	repository repository.CardRepositoryInterface
	validate   *validator.Validate
}

func NewCreateCardHandler(
	repository repository.CardRepositoryInterface,
	validate *validator.Validate,
) *CardCreateHandler {
	return &CardCreateHandler{
		repository: repository,
		validate:   validate,
	}
}

func (s *CardCreateHandler) Handle(ctx context.Context, cmd CreateCardCmd) (*entityPkg.Card, error) {
	err := s.validate.Struct(&cmd)

	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(cmd.UserId)

	if err != nil {
		return nil, err
	}

	entity := entityPkg.NewCard(
		userId,
		cmd.Answer,
		cmd.Question,
		entityPkg.FileType(cmd.FileType),
		cmd.FileId,
	)

	err = s.repository.Create(ctx, entity)

	if err != nil {
		return nil, fmt.Errorf("create new card: %w", err)
	}

	return entity, nil
}
