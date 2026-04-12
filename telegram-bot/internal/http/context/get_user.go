package context

import (
	"errors"
	"github.com/gin-gonic/gin"
	"telegram-bot/internal/domain/entity"
)

type UserKey = string

const key UserKey = "user_key"

var ErrUnableToGetUser = errors.New("unable to get user from context")

func SetUser(c *gin.Context, user *entity.User) {
	c.Set(key, user)
}

func GetUser(c *gin.Context) (*entity.User, error) {
	if val, ok := c.Get(key); ok {
		if user, ok := val.(*entity.User); ok {
			return user, nil
		}
	}

	return nil, ErrUnableToGetUser
}
