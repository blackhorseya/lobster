package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/pb"
	"github.com/google/wire"
)

// IRepo declare okr repo service function
type IRepo interface {
	// QueryByID serve caller to query a objective by id from database
	QueryByID(ctx contextx.Contextx, id string) (*pb.Objective, error)

	// Create serve caller to create a objective to database
	Create(ctx contextx.Contextx, created *pb.Objective) (*pb.Objective, error)

	// List serve caller to list all objectives from database
	List(ctx contextx.Contextx, offset, limit int) ([]*pb.Objective, error)

	// Count serve caller to count all objectives from database
	Count(ctx contextx.Contextx) (int, error)

	// Update serve caller to update a objective to database
	Update(ctx contextx.Contextx, updated *pb.Objective) (*pb.Objective, error)

	// Delete serve caller to delete a objective by id from database
	Delete(ctx contextx.Contextx, id string) (int, error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
