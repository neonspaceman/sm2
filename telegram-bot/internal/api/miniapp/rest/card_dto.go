package rest

type CardResponse struct {
	Id string `json:"id"`
}

type CardsResponse struct {
	Cards     []CardResponse `json:"cards"`
	HasNext   bool           `json:"has_next"`
	EndCursor string         `json:"end_cursor"`
}

type GetCardsQuery struct {
	After string `form:"after" validate:"omitempty,uuid"`
}

type CreateCardRequest struct {
	Question string `json:"question" validate:"required"`
	Answer   string `json:"answer" validate:"required"`
}
