package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/todo"
	"github.com/google/wire"
)

// IRepo declare todo repo service function
type IRepo interface {
	// Create serve caller to create a task to database
	Create(ctx contextx.Contextx, task *todo.Task) (*todo.Task, error)

	// List serve caller to list all tasks from database
	List(ctx contextx.Contextx, offset, limit int) ([]*todo.Task, error)

	// Count serve caller to count all tasks from database
	Count(ctx contextx.Contextx) (int, error)

	// Update serve caller to update a task to database
	Update(ctx contextx.Contextx, updated *todo.Task) (*todo.Task, error)

	// Delete serve caller to delete a task by id from database
	Delete(ctx contextx.Contextx, id string) (int, error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet()
