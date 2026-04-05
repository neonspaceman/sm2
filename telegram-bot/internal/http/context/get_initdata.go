package context

import (
	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type InitDataType = string

const key InitDataType = "initdata"

func SetInitData(c *gin.Context, data *initdata.InitData) {
	c.Set(key, data)
}

func MustGetInitData(c *gin.Context) *initdata.InitData {
	if val, ok := c.Get(key); ok {
		if user, ok := val.(*initdata.InitData); ok {
			return user
		}
	}

	panic("Unable to get init data from context")
}
