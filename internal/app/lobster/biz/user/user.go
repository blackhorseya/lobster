package user

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user/repo"
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/user"
	"github.com/google/wire"
)

// IBiz declare user biz service function
type IBiz interface {
	// GetInfoByID serve caller use userID to get UserInfo
	GetInfoByID(ctx contextx.Contextx, id string) (info *user.Profile, err error)

	// GetInfoByEmail serve caller user userEmail to get UserInfo
	GetInfoByEmail(ctx contextx.Contextx, email string) (info *user.Profile, err error)

	// GetInfoByAccessToken serve caller use user's AccessToken to get userInfo
	GetInfoByAccessToken(ctx contextx.Contextx, token string) (info *user.Profile, err error)

	// Signup serve user to setup a new account
	Signup(ctx contextx.Contextx, email, token string) (info *user.Profile, err error)

	// Login serve user to login with userEmail, and user Password
	Login(ctx contextx.Contextx, email, token string) (info *user.Profile, err error)

	// Logout serve user to logout application which he/she wants
	Logout(ctx contextx.Contextx, user *user.Profile) (err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
