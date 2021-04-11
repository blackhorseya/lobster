package result

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/result/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/okr"
	"github.com/google/wire"
)

// IBiz declare key result biz service function
type IBiz interface {
	// List serve get key result list by page and size
	List(ctx contextx.Contextx, page, size int) (krs []*okr.Result, err error)

	// GetByID serve caller use key result to get key result by id
	GetByID(ctx contextx.Contextx, id string) (kr *okr.Result, err error)

	// GetByGoalID serve caller use goal id to get key results
	GetByGoalID(ctx contextx.Contextx, id string) (krs []*okr.Result, err error)

	// LinkToGoal serve caller link a key result to goal via id
	LinkToGoal(ctx contextx.Contextx, created *okr.Result) (kr *okr.Result, err error)

	// ModifyTitle serve caller to modify title of result
	ModifyTitle(ctx contextx.Contextx, id, title string) (result *okr.Result, err error)

	// Delete serve caller to delete a key result by id
	Delete(ctx contextx.Contextx, id string) (err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	NewImpl,
	repo.ProviderSet,
)
