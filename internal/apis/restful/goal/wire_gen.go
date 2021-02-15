// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package goal

import (
	"github.com/blackhorseya/lobster/internal/biz/goal"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateIHandler(biz goal.IBiz) (IHandler, error) {
	iHandler := NewImpl(biz)
	return iHandler, nil
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl)