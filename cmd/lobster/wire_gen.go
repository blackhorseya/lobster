// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/apis"
	goal2 "github.com/blackhorseya/lobster/internal/app/lobster/apis/goal"
	health2 "github.com/blackhorseya/lobster/internal/app/lobster/apis/health"
	result2 "github.com/blackhorseya/lobster/internal/app/lobster/apis/result"
	task2 "github.com/blackhorseya/lobster/internal/app/lobster/apis/task"
	user2 "github.com/blackhorseya/lobster/internal/app/lobster/apis/user"
	"github.com/blackhorseya/lobster/internal/app/lobster/biz"
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/goal"
	repo3 "github.com/blackhorseya/lobster/internal/app/lobster/biz/goal/repo"
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/health"
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/health/repo"
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/result"
	repo4 "github.com/blackhorseya/lobster/internal/app/lobster/biz/result/repo"
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/task"
	repo2 "github.com/blackhorseya/lobster/internal/app/lobster/biz/task/repo"
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user"
	repo5 "github.com/blackhorseya/lobster/internal/app/lobster/biz/user/repo"
	"github.com/blackhorseya/lobster/internal/pkg/app"
	"github.com/blackhorseya/lobster/internal/pkg/config"
	"github.com/blackhorseya/lobster/internal/pkg/databases"
	"github.com/blackhorseya/lobster/internal/pkg/log"
	"github.com/blackhorseya/lobster/internal/pkg/transports/http"
	"github.com/google/wire"
)

// Injectors from wire.go:

// CreateInjector serve caller to create an injector
func CreateInjector(path2 string) (*app.Injector, error) {
	viper, err := config.New(path2)
	if err != nil {
		return nil, err
	}
	options, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logOptions, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(logOptions)
	if err != nil {
		return nil, err
	}
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
	taskIBiz := task.NewImpl(repoIRepo)
	taskIHandler := task2.NewImpl(taskIBiz)
	iRepo2 := repo3.NewImpl(db)
	goalIBiz := goal.NewImpl(iRepo2)
	goalIHandler := goal2.NewImpl(goalIBiz)
	iRepo3 := repo4.NewImpl(db)
	resultIBiz := result.NewImpl(iRepo3)
	resultIHandler := result2.NewImpl(resultIBiz)
	iRepo4 := repo5.NewImpl(db)
	userIBiz := user.NewImpl(iRepo4)
	userIHandler := user2.NewImpl(userIBiz)
	initHandlers := apis.CreateInitHandlerFn(iHandler, taskIHandler, goalIHandler, resultIHandler, userIHandler)
	engine := http.NewRouter(options, logger, initHandlers)
	injector := app.NewInjector(engine, configConfig)
	return injector, nil
}

// wire.go:

var providerSet = wire.NewSet(app.ProviderSet, log.ProviderSet, config.ProviderSet, http.ProviderSet, databases.ProviderSet, apis.ProviderSet, biz.ProviderSet)
