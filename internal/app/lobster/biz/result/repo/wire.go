// +build wireinject

package repo

import (
	config2 "github.com/blackhorseya/lobster/internal/pkg/entity/config"
	databases2 "github.com/blackhorseya/lobster/internal/pkg/infra/databases"
	log2 "github.com/blackhorseya/lobster/internal/pkg/infra/log"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	log2.ProviderSet,
	config2.ProviderSet,
	databases2.ProviderSet,
	NewImpl,
)

// CreateRepo serve caller to create an IRepo
func CreateRepo(path string) (IRepo, error) {
	panic(wire.Build(testProviderSet))
}
