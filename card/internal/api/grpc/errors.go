package grpc

import (
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func NewErrCardNotFound(cardId uuid.UUID) error {
	st := status.New(codes.NotFound, "card not found")

	stWithDetails, err := st.WithDetails(&errdetails.ErrorInfo{
		Reason: "ERR_CARD_NOT_FOUND",
		Metadata: map[string]string{
			"card_id": cardId.String(),
		},
	})

	if err != nil {
		return fmt.Errorf("attach details: %w", err)
	}

	return stWithDetails.Err()
}

func NewErrReviewPeriodNotStart(due time.Time) error {
	st := status.New(codes.FailedPrecondition, "review period not start")

	stWithDetails, err := st.WithDetails(&errdetails.ErrorInfo{
		Reason: "ERR_REVIEW_PERIOD_NOT_START",
		Metadata: map[string]string{
			"due": due.String(),
		},
	})

	if err != nil {
		return fmt.Errorf("attach details: %w", err)
	}

	return stWithDetails.Err()
}
