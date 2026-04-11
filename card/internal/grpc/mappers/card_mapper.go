package mappers

import (
	domain_card "card/internal/domain/card"
	api_card "card/pkg/api/card"
)

func ToCard(e *domain_card.Card) *api_card.Card {
	return &api_card.Card{
		Id:       e.Id.String(),
		Question: e.Question,
		Answer:   e.Answer,
		FileType: api_card.FileType(api_card.FileType_value[string(e.FileType)]),
		FileId:   e.FileId,
	}
}
