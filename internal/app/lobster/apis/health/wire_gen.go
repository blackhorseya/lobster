// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package health

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/health"
	"github.com/google/wire"
)

// Injectors from wire.go:

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(biz health.IBiz) (IHandler, error) {
	iHandler := NewImpl(biz)
	return iHandler, nil
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl)