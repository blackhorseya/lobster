package middlewares

import (
	"net/http"

	"github.com/blackhorseya/lobster/internal/pkg/entities/errors"
	"github.com/gin-gonic/gin"
)

// ResponseMiddleware serve caller to format api response
func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if c.Errors.Last() == nil {
				return
			}

			err := c.Errors.Last()
			c.Errors = c.Errors[:0]

			switch err.Err.(type) {
			case *errors.APPError:
				appError := err.Err.(*errors.APPError)
				c.AbortWithStatusJSON(appError.Status, appError)
				break
			default:
				c.AbortWithStatus(http.StatusInternalServerError)
				break
			}
		}()

		c.Next()

	}
}
