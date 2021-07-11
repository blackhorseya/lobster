// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package user

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user/repo"
	"github.com/blackhorseya/lobster/internal/pkg/infra/token"
	"github.com/bwmarrin/snowflake"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// Injectors from wire.go:

// CreateIBiz serve caller to create an IBiz
func CreateIBiz(logger *zap.Logger, repo2 repo.IRepo, node *snowflake.Node, token2 *token.Factory) (IBiz, error) {
	iBiz := NewImpl(logger, repo2, node, token2)
	return iBiz, nil
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl)
