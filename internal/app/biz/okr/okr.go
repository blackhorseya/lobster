package okr

import (
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"github.com/google/wire"
)

// IBiz declare okr biz service function
type IBiz interface {
	// SetGoal serve user to create a goal
	SetGoal(ctx contextx.Contextx, goal *okr.Objective) (*okr.Objective, error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet()
