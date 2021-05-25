// +build wireinject

package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/entity/config"
	"github.com/blackhorseya/lobster/internal/pkg/infra/databases"
	"github.com/blackhorseya/lobster/internal/pkg/infra/log"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	databases.ProviderSet,
	config.ProviderSet,
	NewImpl,
)

// CreateIRepo serve caller to create an IRepo
func CreateIRepo(path string) (IRepo, error) {
	panic(wire.Build(testProviderSet))
}
