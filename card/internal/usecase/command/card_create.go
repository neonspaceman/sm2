package command

import (
	entityPkg "card/internal/domain/entity"
	"card/internal/domain/repository"
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type CardCreateCmd struct {
	FrontContent string `validate:"required"`
	BackContent  string `validate:"required"`
}

type CardCreateHandler struct {
	repository repository.CardRepositoryInterface
	validate   *validator.Validate
}

func NewCardCreateHandler(
	repository repository.CardRepositoryInterface,
	validate *validator.Validate,
) *CardCreateHandler {
	return &CardCreateHandler{
		repository: repository,
		validate:   validate,
	}
}

func (s *CardCreateHandler) Handle(ctx context.Context, cmd CardCreateCmd) (entityPkg.CardCard, error) {
	err := s.validate.Struct(&cmd)

	if err != nil {
		return entityPkg.CardCard{}, err
	}

	entity := entityPkg.NewCard(cmd.FrontContent, cmd.BackContent)

	err = s.repository.Create(ctx, entity)

	if err != nil {
		return entityPkg.CardCard{}, fmt.Errorf("cardCreateHandler.Handle: unable to create new card: %w", err)
	}

	return entity, nil
}
