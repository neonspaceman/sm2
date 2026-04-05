package card

import "card/pkg/api/card"

func ToCard(card *card.Card) *Card {
	return &Card{Id: card.Id}
}
