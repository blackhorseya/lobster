package kr

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// IHandler declare key result api handler
type IHandler interface {
	// GetByID serve user to get a key result by id
	GetByID(c *gin.Context)

	// List serve user to list all key results by page and size
	List(c *gin.Context)

	// Create serve user to create a key result for objective
	Create(c *gin.Context)

	// Update serve user to update a key result
	Update(c *gin.Context)

	// Delete serve user to delete a key result by id
	Delete(c *gin.Context)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
