package mappers

import (
	review_domain "card/internal/domain/review"
	card_api "card/pkg/api/card"
	"errors"
	"fmt"
)

var (
	ErrUnknownApiRating    = errors.New("unknown rating type")
	ErrUnknownDomainRating = errors.New("unknown rating representation")
)

func ToRating(rating card_api.Rating) (review_domain.RatingType, error) {
	switch rating {
	case card_api.Rating_AGAIN:
		return review_domain.AgainRating, nil
	case card_api.Rating_HARD:
		return review_domain.HardRating, nil
	case card_api.Rating_GOOD:
		return review_domain.GoodRating, nil
	case card_api.Rating_EASY:
		return review_domain.EasyRating, nil
	default:
		return "", fmt.Errorf("rating type '%d': %w", rating, ErrUnknownApiRating)
	}
}

func FromRating(rating review_domain.RatingType) (card_api.Rating, error) {
	switch rating {
	case review_domain.AgainRating:
		return card_api.Rating_AGAIN, nil
	case review_domain.HardRating:
		return card_api.Rating_HARD, nil
	case review_domain.GoodRating:
		return card_api.Rating_GOOD, nil
	case review_domain.EasyRating:
		return card_api.Rating_EASY, nil
	default:
		return card_api.Rating(0), fmt.Errorf("rating type '%s': %w", rating, ErrUnknownDomainRating)
	}
}
