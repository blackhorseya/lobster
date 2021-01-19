// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/config"
	"github.com/blackhorseya/lobster/internal/pkg/databases"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateRepo(path string) (IRepo, error) {
	configConfig, err := config.NewConfig(path)
	if err != nil {
		return nil, err
	}
	client, err := databases.NewMongoDB(configConfig)
	if err != nil {
		return nil, err
	}
	iRepo := NewImpl(client)
	return iRepo, nil
}

// wire.go:

var testProviderSet = wire.NewSet(config.ProviderSet, databases.ProviderSet, NewImpl)
