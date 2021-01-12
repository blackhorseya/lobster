package okr

import (
	"github.com/google/wire"
)

// IBiz declare okr biz service function
type IBiz interface {
	// SetGoal serve user to create a goal
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet()
