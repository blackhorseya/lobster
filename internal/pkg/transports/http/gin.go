package http

import (
	"github.com/blackhorseya/lobster/internal/pkg/transports/http/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewGinEngine serve caller to create a gin.Engine
func NewGinEngine(init InitHandlers) *gin.Engine {
	// todo: 2021-01-11|14:45|doggy|replace me
	gin.SetMode(gin.DebugMode)

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(middlewares.ContextMiddleware())
	r.Use(middlewares.LoggerMiddleware())

	init(r)

	return r
}
