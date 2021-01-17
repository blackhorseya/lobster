package okr

import (
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"github.com/google/wire"
)

// IBiz declare okr biz service function
type IBiz interface {
	// SetObjective serve user to create a objective
	SetObjective(ctx contextx.Contextx, obj *okr.Objective) (*okr.Objective, error)

	// ListObjectives serve user to list all objectives
	ListObjectives(ctx contextx.Contextx, page, size int) ([]*okr.Objective, error)

	// CountObjective serve user to count all objectives
	CountObjective(ctx contextx.Contextx) (int, error)

	// UpdateObjective serve user to update a objective
	UpdateObjective(ctx contextx.Contextx, updated *okr.Objective) (*okr.Objective, error)

	// DeleteObjective serve user to delete a objective by id
	DeleteObjective(ctx contextx.Contextx, id string) error

	// SetResult serve user to create a key result for objective
	SetResult(ctx contextx.Contextx, id string, result *okr.KeyResult) (*okr.KeyResult, error)

	// UpdateResult serve user to update a key result
	UpdateResult(ctx contextx.Contextx, id string, updated *okr.KeyResult) (*okr.KeyResult, error)

	// DeleteResult serve user to delete a key result by id
	DeleteResult(ctx contextx.Contextx, id string) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet()
