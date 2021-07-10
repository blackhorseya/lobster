package user

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user/repo"
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/base/encrypt"
	"github.com/blackhorseya/lobster/internal/pkg/entity/er"
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
	ret, err := i.repo.GetByID(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrGetUserByID.Error(), zap.Int64("id", id))
		return nil, er.ErrGetUserByID
	}
	if ret == nil {
		i.logger.Error(er.ErrUserNotExists.Error(), zap.Int64("id", id))
		return nil, er.ErrUserNotExists
	}

	return ret, nil
}

func (i *impl) GetByToken(ctx contextx.Contextx, token string) (info *user.Profile, err error) {
	if len(token) == 0 {
		i.logger.Error(er.ErrMissingToken.Error())
		return nil, er.ErrMissingToken
	}

	ret, err := i.repo.GetByToken(ctx, token)
	if err != nil {
		i.logger.Error(er.ErrGetUserByToken.Error(), zap.String("token", token))
		return nil, er.ErrGetUserByToken
	}
	if ret == nil {
		i.logger.Error(er.ErrUserNotExists.Error(), zap.String("token", token))
		return nil, er.ErrUserNotExists
	}

	return ret, nil
}

func (i *impl) Signup(ctx contextx.Contextx, email, password string) (info *user.Profile, err error) {
	if len(email) == 0 {
		i.logger.Error(er.ErrEmptyEmail.Error())
		return nil, er.ErrEmptyEmail
	}

	if len(password) == 0 {
		i.logger.Error(er.ErrEmptyPassword.Error())
		return nil, er.ErrEmptyPassword
	}

	exists, err := i.repo.GetByEmail(ctx, email)
	if err != nil {
		i.logger.Error(er.ErrGetUserByEmail.Error(), zap.String("email", email))
		return nil, er.ErrGetUserByEmail
	}
	if exists != nil {
		i.logger.Error(er.ErrUserEmailExists.Error(), zap.String("email", email))
		return nil, er.ErrUserEmailExists
	}

	salt, err := encrypt.HashAndSalt(password)
	if err != nil {
		return nil, er.ErrEncryptPassword
	}
	ret, err := i.repo.Register(ctx, email, salt)
	if err != nil {
		i.logger.Error(er.ErrSignup.Error(), zap.String("email", email))
		return nil, er.ErrSignup
	}

	return ret, nil
}

func (i *impl) Login(ctx contextx.Contextx, email, password string) (info *user.Profile, err error) {
	// todo: 2021-07-11|07:03|Sean|implement me
	panic("implement me")
}
