package api

import (
	"github.com/blackhorseya/lobster/internal/app/api/health"
	"github.com/blackhorseya/lobster/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func CreateInitHandlerFn(health health.IHandler) http.InitHandlers {
	return func(r *gin.Engine) {
		api := r.Group("/api")
		{
			api.GET("readiness", health.Readiness)
			api.GET("liveness", health.Liveness)
		}
	}
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	CreateInitHandlerFn,
)
