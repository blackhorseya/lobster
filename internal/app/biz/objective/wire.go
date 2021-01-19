// +build wireinject

package objective

import (
	"github.com/blackhorseya/lobster/internal/app/biz/objective/repo"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIBiz serve caller to create an IBiz
func CreateIBiz(repo repo.IRepo) (IBiz, error) {
	panic(wire.Build(testProviderSet))
}
