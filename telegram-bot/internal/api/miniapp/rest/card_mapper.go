package rest

import "telegram-bot/internal/client/card"

func ToCardResponse(dto *card.Card) CardResponse {
	return CardResponse{
		Id: dto.Id,
	}
}
