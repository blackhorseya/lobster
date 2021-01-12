package api

import (
	// import swagger docs
	_ "github.com/blackhorseya/lobster/internal/app/api/docs"
	"github.com/blackhorseya/lobster/internal/app/api/health"
	"github.com/blackhorseya/lobster/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CreateInitHandlerFn serve caller to create init handler
func CreateInitHandlerFn(health health.IHandler) http.InitHandlers {
	return func(r *gin.Engine) {
		api := r.Group("/api")
		{
			api.GET("readiness", health.Readiness)
			api.GET("liveness", health.Liveness)
		}

		if mode := gin.Mode(); mode != gin.ReleaseMode {
			r.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	CreateInitHandlerFn,
)
