package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/samber/lo"
	"net/http"
	"telegram-bot/internal/client/card"
	"telegram-bot/internal/consts"
	"telegram-bot/internal/http/context"
	"telegram-bot/internal/usercase/command"
	"telegram-bot/internal/usercase/query"
)

type CardHandler struct {
	userFirstOrCreateHandler *command.UserFirstOrCreateHandler
	getCardByUserIdHandler   *query.GetCardByUserIdHandler
	validator                *validator.Validate
}

func NewCardHandler(
	userFirstOrCreateHandler *command.UserFirstOrCreateHandler,
	getCardByUserIdHandler *query.GetCardByUserIdHandler,
	validator *validator.Validate,
) *CardHandler {
	return &CardHandler{
		userFirstOrCreateHandler: userFirstOrCreateHandler,
		getCardByUserIdHandler:   getCardByUserIdHandler,
		validator:                validator,
	}
}

func (h *CardHandler) RegisterRoutes(router *gin.Engine) {
	group := router.Group("/cards")
	group.GET("/", h.GetAllCards)
}

func (h *CardHandler) GetAllCards(ctx *gin.Context) {
	var req GetCardsQuery
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.Error(ErrBadRequest)
		return
	}

	err = h.validator.Struct(req)
	if err != nil {
		ctx.Error(err)
		return
	}

	user := context.MustGetUser(ctx)

	cards, hasNext, endCursor, err := h.getCardByUserIdHandler.Handle(ctx, query.GetCardsByUserIdQuery{
		User:  user,
		Limit: consts.GetCardsLimit,
		After: req.After,
	})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, Response[CardsResponse]{
		Data: CardsResponse{
			Cards: lo.Map(cards, func(item *card.Card, i int) CardResponse {
				return CardResponse{
					Id: item.Id,
				}
			}),
			HasNext:   hasNext,
			EndCursor: endCursor,
		},
	})
}
