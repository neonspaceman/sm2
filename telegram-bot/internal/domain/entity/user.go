package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID `db:"id"`
	ChatId    int64     `db:"chat_id"`
	FirstName string    `db:"first_name"`
	CreatedAt time.Time `db:"created_at"`
}

func NewUser(chatId int64, firstName string) *User {
	return &User{
		Id:        uuid.New(),
		ChatId:    chatId,
		FirstName: firstName,
		CreatedAt: time.Now(),
	}
}
