// +build wireinject

package main

import (
	"github.com/blackhorseya/lobster/internal/app/lobster"
	"github.com/blackhorseya/lobster/internal/app/lobster/apis"
	"github.com/blackhorseya/lobster/internal/app/lobster/biz"
	"github.com/blackhorseya/lobster/internal/pkg/app"
	"github.com/blackhorseya/lobster/internal/pkg/entity/config"
	"github.com/blackhorseya/lobster/internal/pkg/infra/databases"
	"github.com/blackhorseya/lobster/internal/pkg/infra/idgen"
	"github.com/blackhorseya/lobster/internal/pkg/infra/log"
	"github.com/blackhorseya/lobster/internal/pkg/infra/token"
	"github.com/blackhorseya/lobster/internal/pkg/infra/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	lobster.ProviderSet,
	log.ProviderSet,
	idgen.ProviderSet,
	config.ProviderSet,
	http.ProviderSet,
	databases.ProviderSet,
	token.ProviderSet,
	apis.ProviderSet,
	biz.ProviderSet,
)

// CreateApp serve caller to create an injector
func CreateApp(path string, nodeID int64) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
