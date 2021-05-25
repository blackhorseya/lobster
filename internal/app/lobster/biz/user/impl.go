package user

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user/repo"
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/base/encrypt"
	"github.com/blackhorseya/lobster/internal/pkg/entity/er"
	"github.com/blackhorseya/lobster/internal/pkg/entity/user"
	"github.com/blackhorseya/lobster/internal/pkg/infra/token"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type impl struct {
	logger       *zap.Logger
	repo         repo.IRepo
	tokenFactory *token.Factory
}

// NewImpl serve caller to create an IBiz
func NewImpl(logger *zap.Logger, repo repo.IRepo, tokenFactory *token.Factory) IBiz {
	return &impl{
		logger:       logger.With(zap.String("type", "UserBiz")),
		repo:         repo,
		tokenFactory: tokenFactory,
	}
}

func (i *impl) Signup(ctx contextx.Contextx, email, password string) (info *user.Profile, err error) {
	if len(email) == 0 {
		i.logger.Error(er.ErrMissingEmail.Error())
		return nil, er.ErrMissingEmail
	}

	if len(password) == 0 {
		i.logger.Error(er.ErrMissingPassword.Error())
		return nil, er.ErrMissingPassword
	}

	exists, err := i.repo.GetByEmail(ctx, email)
	if err != nil {
		i.logger.Error(er.ErrGetUserByEmail.Error(), zap.Error(err), zap.String("email", email))
		return nil, er.ErrGetUserByEmail
	}
	if exists != nil {
		i.logger.Error(er.ErrEmailExists.Error(), zap.String("email", email))
		return nil, er.ErrEmailExists
	}

	enc, err := encrypt.HashAndSalt(password)
	if err != nil {
		i.logger.Error(er.ErrEncryptPassword.Error(), zap.Error(err))
		return nil, er.ErrEncryptPassword
	}
	ret, err := i.repo.Register(ctx, email, enc)
	if err != nil {
		i.logger.Error(er.ErrSignup.Error(), zap.Error(err), zap.String("email", email))
		return nil, er.ErrSignup
	}

	return ret, nil
}

func (i *impl) Login(ctx contextx.Contextx, email, password string) (info *user.Profile, err error) {
	if len(email) == 0 {
		i.logger.Error(er.ErrMissingEmail.Error())
		return nil, er.ErrMissingEmail
	}

	if len(password) == 0 {
		i.logger.Error(er.ErrMissingPassword.Error())
		return nil, er.ErrMissingPassword
	}

	exists, err := i.repo.GetByEmail(ctx, email)
	if err != nil {
		i.logger.Error(er.ErrGetUserByEmail.Error(), zap.Error(err), zap.String("email", email))
		return nil, er.ErrGetUserByEmail
	}
	if exists == nil {
		i.logger.Error(er.ErrUserNotExists.Error(), zap.String("email", email))
		return nil, er.ErrUserNotExists
	}

	err = bcrypt.CompareHashAndPassword([]byte(exists.Password), []byte(password))
	if err != nil {
		i.logger.Error(er.ErrIncorrectPassword.Error(), zap.Error(err), zap.String("email", email))
		return nil, er.ErrIncorrectPassword
	}

	t, err := i.tokenFactory.NewToken(exists.ID, exists.Email)
	if err != nil {
		i.logger.Error(er.ErrNewToken.Error(), zap.Error(err), zap.String("email", email))
		return nil, er.ErrNewToken
	}
	exists.AccessToken = t

	ret, err := i.repo.UpdateToken(ctx, exists)
	if err != nil {
		i.logger.Error(er.ErrLogin.Error(), zap.Error(err), zap.String("email", email))
		return nil, er.ErrLogin
	}

	return ret, nil
}
