// +build wireinject

package main

import (
	"github.com/blackhorseya/lobster/internal/app/lobster"
	"github.com/blackhorseya/lobster/internal/app/lobster/apis"
	"github.com/blackhorseya/lobster/internal/app/lobster/biz"
	"github.com/blackhorseya/lobster/internal/pkg/app"
	config2 "github.com/blackhorseya/lobster/internal/pkg/entity/config"
	databases2 "github.com/blackhorseya/lobster/internal/pkg/infra/databases"
	log2 "github.com/blackhorseya/lobster/internal/pkg/infra/log"
	http2 "github.com/blackhorseya/lobster/internal/pkg/infra/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	lobster.ProviderSet,
	log2.ProviderSet,
	config2.ProviderSet,
	http2.ProviderSet,
	databases2.ProviderSet,
	apis.ProviderSet,
	biz.ProviderSet,
)

// CreateApp serve caller to create an injector
func CreateApp(path string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
