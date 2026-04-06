package context

import (
	"github.com/gin-gonic/gin"
	"telegram-bot/internal/domain/entity"
)

type UserKey = string

const key UserKey = "user_key"

func SetUser(c *gin.Context, user *entity.User) {
	c.Set(key, user)
}

func MustGetUser(c *gin.Context) *entity.User {
	if val, ok := c.Get(key); ok {
		if user, ok := val.(*entity.User); ok {
			return user
		}
	}

	panic("Unable to get user from context")
}
