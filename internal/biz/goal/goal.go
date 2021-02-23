package goal

import (
	"github.com/blackhorseya/lobster/internal/biz/goal/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/pb"
	"github.com/google/wire"
)

// IBiz declare objective biz service function
type IBiz interface {
	// Create serve user to create a objective
	Create(ctx contextx.Contextx, obj *pb.Objective) (*pb.Objective, error)

	// GetByID serve user to get a objective by id
	GetByID(ctx contextx.Contextx, id string) (*pb.Objective, error)

	// List serve user to list all objectives
	List(ctx contextx.Contextx, page, size int) ([]*pb.Objective, error)

	// Count serve user to count all objectives
	Count(ctx contextx.Contextx) (int, error)

	// ModifyTitle serve user to modify title of task
	ModifyTitle(ctx contextx.Contextx, id, title string) (obj *pb.Objective, err error)

	// Delete serve user to delete a objective by id
	Delete(ctx contextx.Contextx, id string) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
