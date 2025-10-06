package entity

import (
	"github.com/google/uuid"
	"time"
)

type CardCard struct {
	id           uuid.UUID `db:"id"`
	frontContent string    `db:"front_content"`
	backContent  string    `db:"back_content"`
	createdAt    time.Time `db:"created_at"`
	updatedAt    time.Time `db:"updated_at"`
}

func NewCard(Title string, Description string) CardCard {
	now := time.Now()

	return CardCard{
		id:           uuid.New(),
		frontContent: Title,
		backContent:  Description,
		createdAt:    now,
		updatedAt:    now,
	}
}

func (e *CardCard) Id() uuid.UUID {
	return e.id
}

func (e *CardCard) FrontContent() string {
	return e.frontContent
}

func (e *CardCard) BackContent() string {
	return e.backContent
}

func (e *CardCard) CreatedAt() time.Time {
	return e.createdAt
}

func (e *CardCard) UpdatedAt() time.Time {
	return e.updatedAt
}
