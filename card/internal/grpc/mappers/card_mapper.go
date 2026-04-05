package mappers

import (
	"card/internal/domain/entity"
	"card/pkg/api/card"
)

func ToCard(e *entity.Card) *card.Card {
	return &card.Card{
		Id:       e.Id.String(),
		Question: e.Question,
		Answer:   e.Answer,
		FileType: card.FileType(card.FileType_value[string(e.FileType)]),
		FileId:   e.FileId,
	}
}
