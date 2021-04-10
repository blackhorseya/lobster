// +build wireinject

package task

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/task"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(logger *zap.Logger, biz task.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
