package task

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/task/repo"
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/task"
	"github.com/google/wire"
)

// IBiz declare todo biz service function
type IBiz interface {
	// GetByID serve user to get a task by id
	GetByID(ctx contextx.Contextx, id int64) (*task.Task, error)

	// List serve user to list all tasks
	List(ctx contextx.Contextx, page, size int) ([]*task.Task, error)

	// Count serve user to count all tasks
	Count(ctx contextx.Contextx) (int, error)

	// Create serve user to create a task
	Create(ctx contextx.Contextx, title string) (*task.Task, error)

	// ModifyTitle serve user to modify title of task
	ModifyTitle(ctx contextx.Contextx, id int64, title string) (t *task.Task, err error)

	// UpdateStatus serve user to update status of task
	UpdateStatus(ctx contextx.Contextx, id int64, status task.Status) (t *task.Task, err error)

	// Delete serve user to delete a task by id
	Delete(ctx contextx.Contextx, id int64) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
