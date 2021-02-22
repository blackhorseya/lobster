package result

import (
	"github.com/blackhorseya/lobster/internal/biz/result/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"github.com/google/wire"
)

// IBiz declare key result biz service function
type IBiz interface {
	// List serve get key result list by page and size
	List(ctx contextx.Contextx, page, size int) (krs []*okr.KeyResult, err error)

	// GetByID serve caller use key result to get key result by id
	GetByID(ctx contextx.Contextx, id string) (kr *okr.KeyResult, err error)

	// GetByGoalID serve caller use goal id to get key results
	GetByGoalID(ctx contextx.Contextx, id string) (krs []*okr.KeyResult, err error)

	// LinkToGoal serve caller link a key result to goal via id
	LinkToGoal(ctx contextx.Contextx, created *okr.KeyResult) (kr *okr.KeyResult, err error)

	// Update serve caller to update a key result
	Update(ctx contextx.Contextx, updated *okr.KeyResult) (kr *okr.KeyResult, err error)

	// Delete serve caller to delete a key result by id
	Delete(ctx contextx.Contextx, id string) (err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	NewImpl,
	repo.ProviderSet,
)
