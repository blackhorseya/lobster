// +build wireinject

package task

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/task/repo"
	"github.com/bwmarrin/snowflake"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIBiz serve caller to create an IBiz
func CreateIBiz(logger *zap.Logger, repo repo.IRepo, node *snowflake.Node) (IBiz, error) {
	panic(wire.Build(testProviderSet))
}
