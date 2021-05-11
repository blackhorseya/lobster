package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/okr"
	"github.com/google/wire"
)

// IRepo declare okr repo service function
type IRepo interface {
	// QueryByID serve caller to query a objective by id from database
	QueryByID(ctx contextx.Contextx, id string) (*okr.Goal, error)

	// Create serve caller to create a objective to database
	Create(ctx contextx.Contextx, created *okr.Goal) (*okr.Goal, error)

	// List serve caller to list all objectives from database
	List(ctx contextx.Contextx, offset, limit int) ([]*okr.Goal, error)

	// Count serve caller to count all objectives from database
	Count(ctx contextx.Contextx) (int, error)

	// Update serve caller to update a objective to database
	Update(ctx contextx.Contextx, updated *okr.Goal) (*okr.Goal, error)

	// Delete serve caller to delete a objective by id from database
	Delete(ctx contextx.Contextx, id string) (int, error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
