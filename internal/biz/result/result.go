package result

import (
	"github.com/blackhorseya/lobster/internal/biz/result/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/pb"
	"github.com/google/wire"
)

// QueryCondition is a struct for caller can applied it to retrieved result list.
type QueryCondition struct {
	ID           string
	Name         string
	UserID       string
	AuthorizerID string
}

// IBiz declare key result biz service function
type IBiz interface {
	// List serve get key result list by page and size
	List(ctx contextx.Contextx, page, size int) (krs []*pb.Result, err error)

	// GetByID serve caller use key result to get key result by id
	GetByID(ctx contextx.Contextx, id string) (kr *pb.Result, err error)

	// GetByGoalID serve caller use goal id to get key results
	GetByGoalID(ctx contextx.Contextx, id string) (krs []*pb.Result, err error)

	// LinkToGoal serve caller link a key result to goal via id
	LinkToGoal(ctx contextx.Contextx, created *pb.Result) (kr *pb.Result, err error)

	// ModifyTitle serve caller to modify title of result
	ModifyTitle(ctx contextx.Contextx, id, title string) (result *pb.Result, err error)

	// Delete serve caller to delete a key result by id
	Delete(ctx contextx.Contextx, id string) (err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	NewImpl,
	repo.ProviderSet,
)
