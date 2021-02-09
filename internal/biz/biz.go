package biz

import (
	"github.com/blackhorseya/lobster/internal/biz/goal"
	"github.com/blackhorseya/lobster/internal/biz/health"
	"github.com/blackhorseya/lobster/internal/biz/kr"
	"github.com/blackhorseya/lobster/internal/biz/task"
	"github.com/google/wire"
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	task.ProviderSet,
	goal.ProviderSet,
	kr.ProviderSet,
)
