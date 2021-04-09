// +build wireinject

package main

import (
	"github.com/blackhorseya/lobster/internal/app/lobster"
	"github.com/blackhorseya/lobster/internal/app/lobster/apis"
	"github.com/blackhorseya/lobster/internal/app/lobster/biz"
	"github.com/blackhorseya/lobster/internal/pkg/app"
	"github.com/blackhorseya/lobster/internal/pkg/config"
	"github.com/blackhorseya/lobster/internal/pkg/databases"
	"github.com/blackhorseya/lobster/internal/pkg/log"
	"github.com/blackhorseya/lobster/internal/pkg/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	lobster.ProviderSet,
	log.ProviderSet,
	config.ProviderSet,
	http.ProviderSet,
	databases.ProviderSet,
	apis.ProviderSet,
	biz.ProviderSet,
)

// CreateApp serve caller to create an injector
func CreateApp(path string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
