package review

import (
	"fmt"
	"time"
)

type ReviewPeriodNotStartError struct {
	ReviewAt time.Time
	Due      time.Time
}

func NewReviewPeriodNotStartError(reviewAt, due time.Time) *ReviewPeriodNotStartError {
	return &ReviewPeriodNotStartError{
		ReviewAt: reviewAt,
		Due:      due,
	}
}

func (e *ReviewPeriodNotStartError) Error() string {
	return fmt.Sprintf("Review date %s is less then need %s", e.ReviewAt.String(), e.Due.String())
}
