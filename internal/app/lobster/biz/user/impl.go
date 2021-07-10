package user

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user/repo"
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/user"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	repo   repo.IRepo
	node   *snowflake.Node
}

// NewImpl serve caller to create an IBiz
func NewImpl(logger *zap.Logger, repo repo.IRepo, node *snowflake.Node) IBiz {
	return &impl{
		logger: logger.With(zap.String("type", "UserBiz")),
		repo:   repo,
		node:   node,
	}
}

func (i *impl) GetByID(ctx contextx.Contextx, id int64) (info *user.Profile, err error) {
	panic("implement me")
}

func (i *impl) GetByToken(ctx contextx.Contextx, token string) (info *user.Profile, err error) {
	panic("implement me")
}

func (i *impl) Signup(ctx contextx.Contextx, email, password string) (info *user.Profile, err error) {
	panic("implement me")
}

func (i *impl) Login(ctx contextx.Contextx, email, password string) (info *user.Profile, err error) {
	panic("implement me")
}
