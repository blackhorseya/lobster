package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type IHandler interface {
	// Me serve caller to get myself information by token
	Me(c *gin.Context)

	// Signup serve caller to create an user by email and password
	Signup(c *gin.Context)

	// Login serve caller to login system by email and password
	Login(c *gin.Context)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
