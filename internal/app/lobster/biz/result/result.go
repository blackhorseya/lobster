package result

import (
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/okr"
	"github.com/google/wire"
)

// IBiz declare result service function
type IBiz interface {
	// GetByID serve caller to given result's id to get result
	GetByID(ctx contextx.Contextx, id int64) (result *okr.Result, err error)

	// List serve caller to list all results
	List(ctx contextx.Contextx, page, size int) (results []*okr.Result, err error)

	// ModifyTitle serve caller to modify result's title
	ModifyTitle(ctx contextx.Contextx, id int64, title string) (result *okr.Result, err error)

	// Delete serve caller to delete a result by id
	Delete(ctx contextx.Contextx, id int64) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet()
