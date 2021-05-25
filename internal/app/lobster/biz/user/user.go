package user

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user/repo"
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/user"
	"github.com/google/wire"
)

// IBiz declare user service function
type IBiz interface {
	// Signup serve caller to register an user
	Signup(ctx contextx.Contextx, email, password string) (info *user.Profile, err error)

	// Login serve caller to login the system
	Login(ctx contextx.Contextx, email, password string) (info *user.Profile, err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
