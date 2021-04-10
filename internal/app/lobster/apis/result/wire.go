// +build wireinject

package result

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/result"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(logger *zap.Logger, biz result.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
