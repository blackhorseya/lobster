// +build wireinject

package health

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/health"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(biz health.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
