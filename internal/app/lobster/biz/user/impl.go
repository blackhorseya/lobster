package user

import (
	"fmt"
	"time"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/user"
	"github.com/google/uuid"
)

var (
	// ErrEmailOrTokenEmpty means email or token is empty
	ErrEmailOrTokenEmpty = fmt.Errorf("email or token is empty")

	// ErrQueryInfoByEmail means get user profile by email is failure
	ErrQueryInfoByEmail = fmt.Errorf("query info by email is failure")

	// ErrUserNotExists means user not exists
	ErrUserNotExists = fmt.Errorf("user not exists")

	// ErrUserLogin means user login failure
	ErrUserLogin = fmt.Errorf("user login failure")

	// ErrUserSignup means user signup failure
	ErrUserSignup = fmt.Errorf("user signup failure")
)

type impl struct {
	repo repo.IRepo
}

// NewImpl serve caller to create an IBiz
func NewImpl(repo repo.IRepo) IBiz {
	return &impl{repo: repo}
}

func (i *impl) GetInfoByID(ctx contextx.Contextx, id string) (info *user.Profile, err error) {
	// todo: 2021-02-28|17:31|doggy|implement me
	panic("implement me")
}

func (i *impl) GetInfoByEmail(ctx contextx.Contextx, email string) (info *user.Profile, err error) {
	logger := ctx.WithField("email", email)

	ret, err := i.repo.QueryInfoByEmail(ctx, email)
	if err != nil {
		logger.WithError(err)
		return nil, err
	}

	return ret, nil
}

func (i *impl) GetInfoByAccessToken(ctx contextx.Contextx, token string) (info *user.Profile, err error) {
	// todo: 2021-02-28|17:31|doggy|implement me
	panic("implement me")
}

func (i *impl) Signup(ctx contextx.Contextx, email, token string) (info *user.Profile, err error) {
	logger := ctx.WithField("email", email).WithField("token", token)

	if len(email) == 0 || len(token) == 0 {
		logger.Error(ErrEmailOrTokenEmpty)
		return nil, ErrEmailOrTokenEmpty
	}

	exist, err := i.repo.QueryInfoByEmail(ctx, email)
	if err != nil {
		logger.WithError(err).Error(ErrQueryInfoByEmail)
		return nil, ErrQueryInfoByEmail
	}
	if exist != nil {
		logger.Error("email is exists")
		return nil, ErrUserSignup
	}

	newUser, err := i.repo.UserRegister(ctx, user.Profile{
		ID:          uuid.New().String(),
		AccessToken: token,
		Email:       email,
		SignupAt:    time.Now().UnixNano(),
	})
	if err != nil {
		logger.WithError(err).Error(ErrUserSignup)
		return nil, ErrUserSignup
	}

	return newUser, nil
}

func (i *impl) Login(ctx contextx.Contextx, email, token string) (info *user.Profile, err error) {
	logger := ctx.WithField("email", email).WithField("token", token)

	if len(email) == 0 || len(token) == 0 {
		logger.Error(ErrEmailOrTokenEmpty)
		return nil, ErrEmailOrTokenEmpty
	}

	exist, err := i.repo.QueryInfoByEmail(ctx, email)
	if err != nil {
		logger.WithError(err).Error(ErrQueryInfoByEmail)
		return nil, ErrQueryInfoByEmail
	}
	if exist == nil {
		logger.Error(ErrUserNotExists)
		return nil, ErrUserNotExists
	}

	if exist.AccessToken != token {
		logger.Error(ErrUserLogin)
		return nil, ErrUserLogin
	}

	return exist, nil
}

func (i *impl) Logout(ctx contextx.Contextx, user *user.Profile) (err error) {
	// todo: 2021-02-28|17:31|doggy|implement me
	panic("implement me")
}
