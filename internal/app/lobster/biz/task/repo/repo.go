package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/todo"
	"github.com/google/wire"
)

// IRepo declare todo repo service function
type IRepo interface {
	// QueryByID serve caller to query a task by id from database
	QueryByID(ctx contextx.Contextx, userID, id int64) (*todo.Task, error)

	// Create serve caller to create a task to database
	Create(ctx contextx.Contextx, task *todo.Task) (*todo.Task, error)

	// List serve caller to list all tasks from database
	List(ctx contextx.Contextx, userID int64, offset, limit int) ([]*todo.Task, error)

	// Count serve caller to count all tasks from database
	Count(ctx contextx.Contextx, userID int64) (int, error)

	// Update serve caller to update a task to database
	Update(ctx contextx.Contextx, updated *todo.Task) (*todo.Task, error)

	// Delete serve caller to delete a task by id from database
	Delete(ctx contextx.Contextx, userID, id int64) (int, error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
