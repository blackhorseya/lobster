package todo

import (
	"github.com/blackhorseya/lobster/internal/app/biz/todo/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/todo"
	"github.com/google/wire"
)

// IBiz declare todo biz service function
type IBiz interface {
	// GetByID serve user to get a task by id
	GetByID(ctx contextx.Contextx, id string) (*todo.Task, error)

	// List serve user to list all tasks
	List(ctx contextx.Contextx, page, size int) ([]*todo.Task, error)

	// Count serve user to count all tasks
	Count(ctx contextx.Contextx) (int, error)

	// Create serve user to create a task
	Create(ctx contextx.Contextx, task *todo.Task) (*todo.Task, error)

	// ChangeTitle serve user to change title of task with title
	ChangeTitle(ctx contextx.Contextx, id, title string) (*todo.Task, error)

	// UpdateStatus serve user to update completed of task
	UpdateStatus(ctx contextx.Contextx, id, completed bool) (*todo.Task, error)

	// Delete serve user to delete a task by id
	Delete(ctx contextx.Contextx, id string) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
