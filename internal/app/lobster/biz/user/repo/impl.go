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


	// todo: 2021-07-11|08:18|Sean|implement me
	panic("implement me")
}

func (i *impl) GetByToken(ctx contextx.Contextx, token string) (info *user.Profile, err error) {
	// todo: 2021-07-11|08:18|Sean|implement me
	panic("implement me")
}

func (i *impl) GetByEmail(ctx contextx.Contextx, email string) (info *user.Profile, err error) {
	// todo: 2021-07-11|08:18|Sean|implement me
	panic("implement me")
}

func (i *impl) Register(ctx contextx.Contextx, newUser *user.Profile) (info *user.Profile, err error) {
	// timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	// defer cancel()
	//
	// stmt := `insert into users values ()`

	// todo: 2021-07-11|08:18|Sean|implement me
	panic("implement me")
}
