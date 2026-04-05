package middleware

import (
	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"net/http"
	"strings"
	"telegram-bot/internal/http/context"
)

const authorizationKey = "tma"

func TelegramAuth(botToken string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authString := c.GetHeader("Authorization")

		if !strings.HasPrefix(authString, authorizationKey+" ") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		authString = strings.SplitN(authString, " ", 2)[1]

		err := initdata.Validate(authString, botToken, 0)

		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
		}

		userData, err := initdata.Parse(authString)

		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
		}

		context.SetInitData(c, &userData)

		c.Next()
	}
}
