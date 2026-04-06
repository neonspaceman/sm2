package middleware

import (
	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"net/http"
	"strings"
	"telegram-bot/internal/http/context"
	"telegram-bot/internal/usercase/command"
)

const authorizationKey = "tma"

func TelegramAuth(botToken string, userFirstOrCreateHandler *command.UserFirstOrCreateHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authString := ctx.GetHeader("Authorization")

		if !strings.HasPrefix(authString, authorizationKey+" ") {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		authString = strings.SplitN(authString, " ", 2)[1]

		err := initdata.Validate(authString, botToken, 0)

		if err != nil {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		userData, err := initdata.Parse(authString)

		if err != nil {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		user, err := userFirstOrCreateHandler.Handle(ctx, userData)
		if err != nil {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		context.SetUser(ctx, user)

		ctx.Next()
	}
}
