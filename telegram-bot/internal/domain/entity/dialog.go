package entity

import (
	"github.com/google/uuid"
	"telegram-bot/internal/domain/types"
	"time"
)

type Dialog struct {
	Id        uuid.UUID          `db:"id"`
	Step      string             `db:"step"`
	Params    types.DialogParams `db:"params"`
	UserId    uuid.UUID          `db:"user_id"`
	CreatedAt time.Time          `db:"created_at"`
	UpdatedAt time.Time          `db:"updated_at"`
}

func NewDialog(step string, params types.DialogParams, userId uuid.UUID) *Dialog {
	now := time.Now()

	return &Dialog{
		Id:        uuid.New(),
		Step:      step,
		Params:    params,
		UserId:    userId,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (d *Dialog) SetStep(step string) *Dialog {
	d.UpdatedAt = time.Now()
	d.Step = step

	return d
}

func (d *Dialog) SetParam(key types.DialogParam, value string) *Dialog {
	d.UpdatedAt = time.Now()
	d.Params[key] = value

	return d
}
