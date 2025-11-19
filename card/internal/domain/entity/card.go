package entity

import (
	"github.com/google/uuid"
	"time"
)

type FileType string

const (
	FileTypeNone     FileType = "NONE"
	FileTypePhoto    FileType = "PHOTO"
	FileTypeDocument FileType = "DOCUMENT"
)

type Card struct {
	Id        uuid.UUID `db:"id"`
	UserId    uuid.UUID `db:"user_id"`
	Question  string    `db:"question"`
	Answer    string    `db:"answer"`
	FileType  FileType  `db:"file_type"`
	FileId    string    `db:"file_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewCard(UserId uuid.UUID, Question, Answer string, FileType FileType, FileId string) *Card {
	now := time.Now()

	return &Card{
		Id:        uuid.New(),
		UserId:    UserId,
		Question:  Question,
		Answer:    Answer,
		FileType:  FileType,
		FileId:    FileId,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
