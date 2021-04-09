// +build wireinject

package task

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/task/repo"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIBiz serve caller to create an IBiz
func CreateIBiz(repo repo.IRepo) (IBiz, error) {
	panic(wire.Build(testProviderSet))
}
