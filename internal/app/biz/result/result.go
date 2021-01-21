package result

import (
	"github.com/blackhorseya/lobster/internal/app/biz/result/repo"
	"github.com/google/wire"
)

// IBiz declare key result biz service function
type IBiz interface {
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
