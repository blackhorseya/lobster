// +build wireinject

package user

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user/repo"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIBiz serve caller to create an IBiz
func CreateIBiz(repo repo.IRepo) (IBiz, error) {
	panic(wire.Build(testProviderSet))
}
