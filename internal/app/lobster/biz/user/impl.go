package user

import (
	"time"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user/repo"
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/base/encrypt"
	"github.com/blackhorseya/lobster/internal/pkg/entity/er"
	"github.com/blackhorseya/lobster/internal/pkg/entity/user"
	"github.com/blackhorseya/lobster/internal/pkg/infra/token"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type impl struct {
	logger *zap.Logger
	repo   repo.IRepo
	node   *snowflake.Node
	token  *token.Factory
}

// NewImpl serve caller to create an IBiz
func NewImpl(logger *zap.Logger, repo repo.IRepo, node *snowflake.Node, token *token.Factory) IBiz {
	return &impl{
		logger: logger.With(zap.String("type", "UserBiz")),
		repo:   repo,
		node:   node,
		token:  token,
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

	ret, err := i.repo.Register(ctx, &user.Profile{
		ID:        i.node.Generate().Int64(),
		Email:     email,
		Password:  salt,
		CreatedAt: time.Now().UnixNano(),
	})
	if err != nil {
		i.logger.Error(er.ErrSignup.Error(), zap.String("email", email))
		return nil, er.ErrSignup
	}

	return ret, nil
}

func (i *impl) Login(ctx contextx.Contextx, email, password string) (info *user.Profile, err error) {
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
	if exists == nil {
		i.logger.Error(er.ErrUserNotExists.Error(), zap.String("email", email))
		return nil, er.ErrUserNotExists
	}

	err = bcrypt.CompareHashAndPassword([]byte(exists.Password), []byte(password))
	if err != nil {
		i.logger.Error(er.ErrIncorrectPassword.Error(), zap.String("email", email))
		return nil, er.ErrIncorrectPassword
	}

	newToken, err := i.token.NewToken(exists.ID, exists.Email)
	if err != nil {
		i.logger.Error(er.ErrNewToken.Error(), zap.Int64("id", exists.ID), zap.String("email", email))
		return nil, er.ErrNewToken
	}
	exists.Token = newToken

	ret, err := i.repo.UpdateToken(ctx, exists)
	if err != nil {
		i.logger.Error(er.ErrUpdateToken.Error(), zap.Int64("id", exists.ID), zap.String("email", email), zap.String("token", newToken))
		return nil, er.ErrUpdateToken
	}

	return ret, nil
}
