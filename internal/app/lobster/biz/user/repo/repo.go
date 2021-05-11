package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/user"
	"github.com/google/wire"
)

// IRepo declare user repo service function
type IRepo interface {
	// QueryInfoByEmail serve caller to given email to get profile
	QueryInfoByEmail(ctx contextx.Contextx, email string) (info *user.Profile, err error)

	// UserRegister serve caller to register a user
	UserRegister(ctx contextx.Contextx, newUser user.Profile) (info *user.Profile, err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
