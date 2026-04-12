package card

import (
	"fmt"
	"github.com/google/uuid"
)

type CardNotFoundError struct {
	CardId uuid.UUID
}

func NewCardNotFoundError(cardId uuid.UUID) *CardNotFoundError {
	return &CardNotFoundError{CardId: cardId}
}

func (e *CardNotFoundError) Error() string {
	return fmt.Sprintf("card '%s' not found", e.CardId.String())
}
