package objective

import (
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
)

// IBiz declare objective biz service function
type IBiz interface {
	// Create serve user to create a objective
	Create(ctx contextx.Contextx, obj *okr.Objective) (*okr.Objective, error)

	// List serve user to list all objectives
	List(ctx contextx.Contextx, page, size int) ([]*okr.Objective, error)

	// Count serve user to count all objectives
	Count(ctx contextx.Contextx) (int, error)

	// Update serve user to update a objective
	Update(ctx contextx.Contextx, updated *okr.Objective) (*okr.Objective, error)

	// Delete serve user to delete a objective by id
	Delete(ctx contextx.Contextx, id string) error
}
