package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/user"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	rw     *sqlx.DB
}

// NewImpl serve caller to create an IRepo
func NewImpl(logger *zap.Logger, rw *sqlx.DB) IRepo {
	return &impl{
		logger: logger.With(zap.String("type", "UserRepo")),
		rw:     rw,
	}
}

func (i *impl) GetByID(ctx contextx.Contextx, id int64) (info *user.Profile, err error) {
	panic("implement me")
}

func (i *impl) GetByToken(ctx contextx.Contextx, token string) (info *user.Profile, err error) {
	panic("implement me")
}

func (i *impl) GetByEmail(ctx contextx.Contextx, email string) (info *user.Profile, err error) {
	panic("implement me")
}

func (i *impl) Register(ctx contextx.Contextx, email, password string) (info *user.Profile, err error) {
	panic("implement me")
}
