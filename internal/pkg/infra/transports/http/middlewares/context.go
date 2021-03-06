package middlewares

import (
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/base/trace"
	"github.com/gin-gonic/gin"
)

// ContextMiddleware serve caller to added Contextx into gin
func ContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := contextx.WithValue(contextx.Background(), "traceID", trace.NewTraceID())
		c.Set("ctx", ctx)

		c.Next()
	}
}
