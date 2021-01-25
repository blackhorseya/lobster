// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/blackhorseya/lobster/internal/app"
	"github.com/blackhorseya/lobster/internal/app/apis"
	health2 "github.com/blackhorseya/lobster/internal/app/apis/health"
	objective2 "github.com/blackhorseya/lobster/internal/app/apis/objective"
	todo2 "github.com/blackhorseya/lobster/internal/app/apis/todo"
	"github.com/blackhorseya/lobster/internal/app/biz"
	"github.com/blackhorseya/lobster/internal/app/biz/health"
	"github.com/blackhorseya/lobster/internal/app/biz/health/repo"
	"github.com/blackhorseya/lobster/internal/app/biz/objective"
	repo3 "github.com/blackhorseya/lobster/internal/app/biz/objective/repo"
	"github.com/blackhorseya/lobster/internal/app/biz/todo"
	repo2 "github.com/blackhorseya/lobster/internal/app/biz/todo/repo"
	"github.com/blackhorseya/lobster/internal/pkg/config"
	"github.com/blackhorseya/lobster/internal/pkg/databases"
	"github.com/blackhorseya/lobster/internal/pkg/transports/http"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateInjector(path2 string) (*app.Injector, error) {
	configConfig, err := config.NewConfig(path2)
	if err != nil {
		return nil, err
	}
	db, err := databases.NewMariaDB(configConfig)
	if err != nil {
		return nil, err
	}
	iRepo := repo.NewImpl(db)
	iBiz := health.NewImpl(iRepo)
	iHandler := health2.NewImpl(iBiz)
	repoIRepo := repo2.NewImpl(db)
	todoIBiz := todo.NewImpl(repoIRepo)
	todoIHandler := todo2.NewImpl(todoIBiz)
	iRepo2 := repo3.NewImpl(db)
	objectiveIBiz := objective.NewImpl(iRepo2)
	objectiveIHandler := objective2.NewImpl(objectiveIBiz)
	initHandlers := apis.CreateInitHandlerFn(iHandler, todoIHandler, objectiveIHandler)
	engine := http.NewGinEngine(configConfig, initHandlers)
	injector := app.NewInjector(engine, configConfig)
	return injector, nil
}

// wire.go:

var providerSet = wire.NewSet(app.ProviderSet, config.ProviderSet, http.ProviderSet, databases.ProviderSet, apis.ProviderSet, biz.ProviderSet)
