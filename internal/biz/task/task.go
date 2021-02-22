package task

import (
	"github.com/blackhorseya/lobster/internal/biz/task/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities"
	"github.com/google/wire"
)

// IBiz declare todo biz service function
type IBiz interface {
	// GetByID serve user to get a task by id
	GetByID(ctx contextx.Contextx, id string) (*task.Task, error)

	// List serve user to list all tasks
	List(ctx contextx.Contextx, page, size int) ([]*task.Task, error)

	// Count serve user to count all tasks
	Count(ctx contextx.Contextx) (int, error)

	// Create serve user to create a task
	Create(ctx contextx.Contextx, task *task.Task) (*task.Task, error)

	// Update serve user to update a task
	Update(ctx contextx.Contextx, updated *task.Task) (*task.Task, error)

	// Delete serve user to delete a task by id
	Delete(ctx contextx.Contextx, id string) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
