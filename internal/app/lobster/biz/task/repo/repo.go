package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/pb"
	"github.com/google/wire"
)

// IRepo declare todo repo service function
type IRepo interface {
	// QueryByID serve caller to query a task by id from database
	QueryByID(ctx contextx.Contextx, id string) (*pb.Task, error)

	// Create serve caller to create a task to database
	Create(ctx contextx.Contextx, task *pb.Task) (*pb.Task, error)

	// List serve caller to list all tasks from database
	List(ctx contextx.Contextx, offset, limit int) ([]*pb.Task, error)

	// Count serve caller to count all tasks from database
	Count(ctx contextx.Contextx) (int, error)

	// Update serve caller to update a task to database
	Update(ctx contextx.Contextx, updated *pb.Task) (*pb.Task, error)

	// Delete serve caller to delete a task by id from database
	Delete(ctx contextx.Contextx, id string) (int, error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)