package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/user"
	"github.com/google/wire"
)

// IRepo declare user repository function
type IRepo interface {
	// GetByEmail serve caller to given email to get an user information
	GetByEmail(ctx contextx.Contextx, email string) (info *user.Profile, err error)

	// Register serve caller to create an user
	Register(ctx contextx.Contextx, email, password string) (info *user.Profile, err error)

	// UpdateToken serve caller to given a jwt then update to database
	UpdateToken(ctx contextx.Contextx, updated *user.Profile) (info *user.Profile, err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
