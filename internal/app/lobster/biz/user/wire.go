// +build wireinject

package user

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user/repo"
	"github.com/blackhorseya/lobster/internal/pkg/entity/config"
	"github.com/blackhorseya/lobster/internal/pkg/infra/log"
	"github.com/blackhorseya/lobster/internal/pkg/infra/token"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	token.ProviderSet,
	NewImpl,
)

// CreateIBiz serve caller to create an IBiz
func CreateIBiz(path string, repo repo.IRepo) (IBiz, error) {
	panic(wire.Build(testProviderSet))
}
