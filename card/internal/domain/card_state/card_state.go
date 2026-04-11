package card_state

import (
	"github.com/google/uuid"
	"time"
)

type CardStateType string

const (
	LearningCardState   CardStateType = "LEARNING"
	ReviewCardState     CardStateType = "REVIEW"
	RelearningCardState CardStateType = "RELEARNING"
)

type CardState struct {
	Id                    uuid.UUID     `db:"id"`
	State                 CardStateType `db:"state"`
	Step                  int           `db:"step"`
	Easy                  float64       `db:"easy"`
	Due                   time.Time     `db:"due"`
	CurrentIntervalInDays int64         `db:"current_interval_in_days"`
	CreatedAt             time.Time     `db:"created_at"`
	UpdatedAt             time.Time     `db:"updated_at"`
}

func NewCardState(id uuid.UUID) *CardState {
	now := time.Now()

	return &CardState{
		Id:        id,
		State:     LearningCardState,
		Due:       now,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (s *CardState) SetState(state CardStateType) *CardState {
	s.State = state

	return s
}

func (s *CardState) SetStep(step int) *CardState {
	s.Step = step

	return s
}

func (s *CardState) SetEasy(easy float64) *CardState {
	s.Easy = easy

	return s
}

func (s *CardState) SetDue(due time.Time) *CardState {
	s.Due = due

	return s
}

func (s *CardState) SetCurrentIntervalInDays(currentIntervalInDays int64) *CardState {
	s.CurrentIntervalInDays = currentIntervalInDays

	return s
}

func (s *CardState) BeforeUpdate() *CardState {
	s.UpdatedAt = time.Now()

	return s
}
