package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// InitHandlers define register handler
type InitHandlers func(r *gin.Engine)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewGinEngine)
