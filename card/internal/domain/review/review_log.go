package review

import (
	"github.com/google/uuid"
	"time"
)

type ReviewLog struct {
	Id        uuid.UUID  `db:"id"`
	CardId    uuid.UUID  `db:"card_id"`
	Rating    RatingType `db:"rating"`
	CreatedAt time.Time  `db:"created_at"`
}

func NewReviewLog(cardId uuid.UUID, rating RatingType, createdAt time.Time) (*ReviewLog, error) {
	id, err := uuid.NewV7()

	if err != nil {
		return nil, err
	}

	return &ReviewLog{
		Id:        id,
		CardId:    cardId,
		Rating:    rating,
		CreatedAt: time.Now(),
	}, nil
}
