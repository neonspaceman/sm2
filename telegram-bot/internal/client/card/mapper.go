package card

import (
	card_api "card/pkg/api/card"
	"fmt"
)

func ToCard(card *card_api.Card) *Card {
	return &Card{Id: card.Id}
}

func ToReviewLog(reviewLog *card_api.ReviewLog) *ReviewLog {
	return &ReviewLog{Id: reviewLog.Id}
}

func FromRating(rating string) (card_api.Rating, error) {
	v, ok := card_api.Rating_value[rating]

	if !ok {
		return 0, fmt.Errorf("parse '%s': %w", rating, ErrUnknownRating)
	}

	return card_api.Rating(v), nil
}
