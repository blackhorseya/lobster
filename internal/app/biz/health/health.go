package health

import (
	"github.com/blackhorseya/lobster/internal/app/biz/health/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/google/wire"
)

// IBiz declare health biz service function
type IBiz interface {
	// Readiness to handle application has been ready
	Readiness(ctx contextx.Contextx) error

	// Liveness to handle application was alive
	Liveness(ctx contextx.Contextx) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
