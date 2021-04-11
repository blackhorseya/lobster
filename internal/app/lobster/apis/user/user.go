package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// IHandler declare todo api handler
type IHandler interface {
	// Signup serve caller to register a user
	Signup(c *gin.Context)

	// Login serve caller to login system
	Login(c *gin.Context)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
