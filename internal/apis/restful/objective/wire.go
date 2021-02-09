// +build wireinject

package objective

import (
	"github.com/blackhorseya/lobster/internal/biz/objective"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(biz objective.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
