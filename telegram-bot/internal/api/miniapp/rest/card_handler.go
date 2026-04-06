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
	userFirstOrCreateHandler *command.FirstOrCreateUserHandler
	getCardByUserIdHandler   *query.GetCardByUserIdHandler
	createCardHandler        *command.CreateCardHandler
	validator                *validator.Validate
}

func NewCardHandler(
	userFirstOrCreateHandler *command.FirstOrCreateUserHandler,
	getCardByUserIdHandler *query.GetCardByUserIdHandler,
	createCardHandler *command.CreateCardHandler,
	validator *validator.Validate,
) *CardHandler {
	return &CardHandler{
		userFirstOrCreateHandler: userFirstOrCreateHandler,
		getCardByUserIdHandler:   getCardByUserIdHandler,
		createCardHandler:        createCardHandler,
		validator:                validator,
	}
}

func (h *CardHandler) RegisterRoutes(router *gin.Engine) {
	group := router.Group("/cards")
	group.GET("/", h.GetAllCards)
	group.POST("/", h.CreateCard)
}

func (h *CardHandler) CreateCard(ctx *gin.Context) {
	var req CreateCardRequest
	err := ctx.ShouldBindJSON(&req)
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

	createdCard, err := h.createCardHandler.Handle(ctx, command.CreateCardCmd{
		User:     user,
		Answer:   req.Answer,
		Question: req.Question,
	})

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, ToCardResponse(createdCard))
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
				return ToCardResponse(item)
			}),
			HasNext:   hasNext,
			EndCursor: endCursor,
		},
	})
}
