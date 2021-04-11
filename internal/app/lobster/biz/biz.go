package biz

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/goal"
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/health"
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/result"
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/task"
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user"
	"github.com/google/wire"
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	task.ProviderSet,
	goal.ProviderSet,
	result.ProviderSet,
	user.ProviderSet,
)
