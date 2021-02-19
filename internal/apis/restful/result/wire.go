// +build wireinject

package result

import (
	"github.com/blackhorseya/lobster/internal/biz/kr"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(biz kr.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
