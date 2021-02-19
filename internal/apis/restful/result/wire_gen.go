// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package result

import (
	"github.com/blackhorseya/lobster/internal/biz/result"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateIHandler(biz result.IBiz) (IHandler, error) {
	iHandler := NewImpl(biz)
	return iHandler, nil
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl)
