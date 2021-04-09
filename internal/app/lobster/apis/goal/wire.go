// +build wireinject

package goal

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/goal"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(biz goal.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
