package card

import "errors"

var (
	ErrUnknownRating        = errors.New("unknown rating type")
	ErrCardNotFound         = errors.New("card not found")
	ErrReviewPeriodNotStart = errors.New("review period not start")
)
