// +build wireinject

package todo

import (
	"github.com/blackhorseya/lobster/internal/biz/todo"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(biz todo.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
