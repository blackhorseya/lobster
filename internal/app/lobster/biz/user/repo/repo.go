package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/entities/user"
	"github.com/blackhorseya/lobster/internal/pkg/utils/contextx"
	"github.com/google/wire"
)

// IRepo declare user repo service function
type IRepo interface {
	// QueryInfoByEmail
	QueryInfoByEmail(ctx contextx.Contextx, email string) (info *user.Profile, err error)

	// UserRegister
	UserRegister(ctx contextx.Contextx, newUser user.Profile) (info *user.Profile, err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
