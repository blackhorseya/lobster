// +build wireinject

package result

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/result"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(biz result.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
