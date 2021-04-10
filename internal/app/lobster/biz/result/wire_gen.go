// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package result

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/result/repo"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func CreateIBiz(logger *zap.Logger, repo2 repo.IRepo) (IBiz, error) {
	iBiz := NewImpl(logger, repo2)
	return iBiz, nil
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl)
