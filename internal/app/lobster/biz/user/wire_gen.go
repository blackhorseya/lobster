// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package user

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user/repo"
	"github.com/blackhorseya/lobster/internal/pkg/entity/config"
	"github.com/blackhorseya/lobster/internal/pkg/infra/log"
	"github.com/blackhorseya/lobster/internal/pkg/infra/token"
	"github.com/google/wire"
)

// Injectors from wire.go:

// CreateIBiz serve caller to create an IBiz
func CreateIBiz(path string, repo2 repo.IRepo) (IBiz, error) {
	viper, err := config.New(path)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	tokenOptions, err := token.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	factory, err := token.New(tokenOptions, logger)
	if err != nil {
		return nil, err
	}
	iBiz := NewImpl(logger, repo2, factory)
	return iBiz, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, token.ProviderSet, NewImpl)
