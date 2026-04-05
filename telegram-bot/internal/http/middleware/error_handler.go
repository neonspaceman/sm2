package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"platform/pkg/logger"
	"telegram-bot/internal/api/miniapp/rest"
)

func ErrorHandler(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			var httpError *rest.HttpError
			if errors.As(err, &httpError) {
				c.AbortWithStatusJSON(httpError.StatusCode, httpError)
				return
			}

			var validationErr validator.ValidationErrors
			if errors.As(err, &validationErr) {
				var validationDetails []gin.H

				for _, v := range validationErr {
					validationDetails = append(validationDetails, gin.H{
						"field":      v.Field(),
						"validation": v.Tag(),
					})
				}

				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"code":    "validation_error",
					"message": "Validation errors",
					"details": gin.H{"errors": validationDetails},
				})
				return
			}

			log.Error("Unhandled error", zap.Error(err))

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    "internal_error",
				"message": "An unexpected error occurred",
			})
		}
	}
}
