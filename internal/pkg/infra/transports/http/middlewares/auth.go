package middlewares

import (
	"strings"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user"
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/er"
	"github.com/gin-gonic/gin"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}

// AuthMiddleware serve caller to extract authorization header value
func AuthMiddleware(userB user.IBiz) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}
		err := c.ShouldBindHeader(&h)
		if err != nil {
			c.Error(er.ErrMissingToken)
			c.Abort()
			return
		}

		idTokenHeader := strings.Split(h.IDToken, "Bearer ")
		if len(idTokenHeader) < 2 {
			c.Error(er.ErrAuthHeaderFormat)
			c.Abort()
			return
		}

		token := idTokenHeader[1]
		ctx := c.MustGet("ctx").(contextx.Contextx)

		info, err := userB.GetByToken(ctx, token)
		if err != nil {
			c.Error(err)
			return
		}

		c.Set("ctx", contextx.WithValue(ctx, "user", info))

		c.Next()
	}
}
