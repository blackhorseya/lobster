package biz

import (
	"github.com/blackhorseya/lobster/internal/biz/health"
	"github.com/blackhorseya/lobster/internal/biz/objective"
	"github.com/blackhorseya/lobster/internal/biz/result"
	"github.com/blackhorseya/lobster/internal/biz/task"
	"github.com/google/wire"
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	task.ProviderSet,
	objective.ProviderSet,
	result.ProviderSet,
)
