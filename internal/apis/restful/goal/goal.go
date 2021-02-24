package goal

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// IHandler declare objective api handler
type IHandler interface {
	// GetByID serve user to get a objective by id
	GetByID(c *gin.Context)

	// List serve user to list all objectives
	List(c *gin.Context)

	// Create serve user to creat a objective
	Create(c *gin.Context)

	// ModifyTitle serve user to modify title of goal
	ModifyTitle(c *gin.Context)

	// Delete serve user to delete a objective by id
	Delete(c *gin.Context)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
