// +build wireinject

package main

import (
	"github.com/blackhorseya/lobster/internal/app"
	"github.com/blackhorseya/lobster/internal/app/apis"
	"github.com/blackhorseya/lobster/internal/app/biz"
	"github.com/blackhorseya/lobster/internal/pkg/config"
	"github.com/blackhorseya/lobster/internal/pkg/databases"
	"github.com/blackhorseya/lobster/internal/pkg/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	app.ProviderSet,
	config.ProviderSet,
	http.ProviderSet,
	databases.ProviderSet,
	apis.ProviderSet,
	biz.ProviderSet,
)

// CreateInjector serve caller to create an injector
func CreateInjector(path string) (*app.Injector, error) {
	panic(wire.Build(providerSet))
}