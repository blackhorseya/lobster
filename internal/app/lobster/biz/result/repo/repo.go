package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/okr"
	"github.com/google/wire"
)

// IRepo declare result repo function
type IRepo interface {
	// GetByID serve caller to get a result by user id and id
	GetByID(ctx contextx.Contextx, userID, id int64) (result *okr.Result, err error)

	// List serve caller to list all results
	List(ctx contextx.Contextx, userID int64, limit, offset int) (results *okr.Result, err error)

	// Create serve caller to create a result
	Create(ctx contextx.Contextx, created *okr.Result) (result *okr.Result, err error)

	// Update serve caller to update a result
	Update(ctx contextx.Contextx, updated *okr.Result) (result *okr.Result, err error)

	// Delete serve caller to delete a result by user id and id
	Delete(ctx contextx.Contextx, userID, id int64) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet()
