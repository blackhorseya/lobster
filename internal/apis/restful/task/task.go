package task

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// IHandler declare todo api handler
type IHandler interface {
	// GetByID serve user to get a task by id
	GetByID(c *gin.Context)

	// List serve user to list all tasks
	List(c *gin.Context)

	// Create serve user to creat a task
	Create(c *gin.Context)

	// ModifyTitle serve user to modify title of task
	ModifyTitle(c *gin.Context)

	// UpdateStatus serve user to update status of task
	UpdateStatus(c *gin.Context)

	// Delete serve user to delete a task by id
	Delete(c *gin.Context)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
