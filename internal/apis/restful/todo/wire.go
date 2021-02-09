// +build wireinject

package todo

import (
	"github.com/blackhorseya/lobster/internal/biz/task"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(biz task.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
