package biz

import (
	"github.com/blackhorseya/lobster/internal/app/biz/okr"
	"github.com/google/wire"
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(okr.ProviderSet)
