package http

import (
	"github.com/blackhorseya/lobster/internal/pkg/config"
	"github.com/blackhorseya/lobster/internal/pkg/transports/http/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewGinEngine serve caller to create a gin.Engine
func NewGinEngine(cfg *config.Config, init InitHandlers) *gin.Engine {
	gin.SetMode(cfg.HTTP.Mode)

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(middlewares.ContextMiddleware())
	r.Use(middlewares.LoggerMiddleware())

	init(r)

	return r
}
