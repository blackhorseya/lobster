package api

import (
	"github.com/blackhorseya/lobster/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func CreateInitHandlerFn() http.InitHandlers {
	return func(r *gin.Engine) {

	}
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(CreateInitHandlerFn)
