// +build wireinject

package goal

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/goal/repo"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIBiz serve caller to create an IBiz
func CreateIBiz(repo repo.IRepo) (IBiz, error) {
	panic(wire.Build(testProviderSet))
}
