package user

import (
	"github.com/blackhorseya/lobster/internal/biz/user/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/pb"
	"github.com/google/wire"
)

// IBiz declare user biz service function
type IBiz interface {
	// GetInfoByID serve caller use userID to get UserInfo
	GetInfoByID(ctx contextx.Contextx, id string) (info *pb.Profile, err error)

	// GetInfoByEmail serve caller user userEmail to get UserInfo
	GetInfoByEmail(ctx contextx.Contextx, email string) (info *pb.Profile, err error)

	// GetInfoByAccessToken serve caller use user's AccessToken to get userInfo
	GetInfoByAccessToken(ctx contextx.Contextx, token string) (info *pb.Profile, err error)

	// Login serve user to login with userEmail, and user Password
	Login(ctx contextx.Contextx, email, token string) (info *pb.Profile, err error)

	// Logout serve user to logout application which he/she wants
	Logout(ctx contextx.Contextx, user *pb.Profile) (err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
