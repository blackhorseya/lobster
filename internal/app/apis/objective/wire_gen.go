// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package objective

import (
	"github.com/blackhorseya/lobster/internal/app/biz/objective"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateIHandler(biz objective.IBiz) (IHandler, error) {
	iHandler := NewImpl(biz)
	return iHandler, nil
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl)
