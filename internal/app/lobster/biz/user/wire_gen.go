// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package user

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user/repo"
	"github.com/google/wire"
)

// Injectors from wire.go:

// CreateIBiz serve caller to create an IBiz
func CreateIBiz(repo2 repo.IRepo) (IBiz, error) {
	iBiz := NewImpl(repo2)
	return iBiz, nil
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl)