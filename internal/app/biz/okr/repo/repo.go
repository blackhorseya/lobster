package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"github.com/google/wire"
)

// IRepo declare okr repo service function
type IRepo interface {
	// QueryObjectiveByID serve caller to query a objective by id from database
	QueryObjectiveByID(ctx contextx.Contextx, id string) (*okr.Objective, error)

	// CreateObjective serve caller to create a objective to database
	CreateObjective(ctx contextx.Contextx, created *okr.Objective) (*okr.Objective, error)

	// ListObjectives serve caller to list all objectives from database
	ListObjectives(ctx contextx.Contextx, offset, limit int) ([]*okr.Objective, error)

	// CountObjective serve caller to count all objectives from database
	CountObjective(ctx contextx.Contextx) (int, error)

	// UpdateObjective serve caller to update a objective to database
	UpdateObjective(ctx contextx.Contextx, updated *okr.Objective) (*okr.Objective, error)

	// DeleteObjective serve caller to delete a objective by id from database
	DeleteObjective(ctx contextx.Contextx, id string) (int, error)
}

// todo: 2021-01-18|21:58|doggy|inject implement
// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet()
