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

type ReviewCardRequest struct {
	CardId string `json:"card_id" validate:"required,uuid"`
	Rating string `json:"rating" validate:"required,oneof=AGAIN HARD GOOD EASY"`
}

type ReviewCardResponse struct {
	ReviewLogId string `json:"review_log_id"`
}
