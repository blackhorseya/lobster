package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/user"
	"github.com/google/wire"
)

// IRepo declare user repo function
type IRepo interface {
	// GetByID serve caller to get user information by id
	GetByID(ctx contextx.Contextx, id int64) (info *user.Profile, err error)

	// GetByToken serve caller to get user information by token
	GetByToken(ctx contextx.Contextx, token string) (info *user.Profile, err error)

	// GetByEmail serve caller to get user information by email
	GetByEmail(ctx contextx.Contextx, email string) (info *user.Profile, err error)

	// Register serve caller to create an account by email and password
	Register(ctx contextx.Contextx, newUser *user.Profile) (info *user.Profile, err error)

	// UpdateToken serve caller to update token by id
	UpdateToken(ctx contextx.Contextx, updated *user.Profile) (info *user.Profile, err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
