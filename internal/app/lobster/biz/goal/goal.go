package goal

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/goal/repo"
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/okr"
	"github.com/google/wire"
)

// IBiz declare objective biz service function
type IBiz interface {
	// Create serve user to create a objective
	Create(ctx contextx.Contextx, obj *okr.Goal) (*okr.Goal, error)

	// GetByID serve user to get a objective by id
	GetByID(ctx contextx.Contextx, id string) (*okr.Goal, error)

	// List serve user to list all objectives
	List(ctx contextx.Contextx, page, size int) ([]*okr.Goal, error)

	// ModifyTitle serve user to modify title of task
	ModifyTitle(ctx contextx.Contextx, id, title string) (obj *okr.Goal, err error)

	// Delete serve user to delete a objective by id
	Delete(ctx contextx.Contextx, id string) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
