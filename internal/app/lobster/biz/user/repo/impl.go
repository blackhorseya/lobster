package repo

import (
	"database/sql"
	"time"

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

func (i *impl) GetByEmail(ctx contextx.Contextx, email string) (info *user.Profile, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	ret := user.Profile{}
	stmt := `select id, email, password, access_token from users where email = ?`
	err = i.rw.GetContext(timeout, &ret, stmt, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &ret, nil
}

func (i *impl) Register(ctx contextx.Contextx, email, password string) (info *user.Profile, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := `insert into users (email, password, access_token) values (:email, :password, :access_token)`
	res, err := i.rw.NamedExecContext(timeout, stmt, &user.Profile{Email: email, Password: password, AccessToken: ""})
	if err != nil {
		return nil, err
	}

	id, _ := res.LastInsertId()
	return &user.Profile{
		ID:          id,
		Email:       email,
		Password:    password,
		AccessToken: "",
	}, nil
}

func (i *impl) UpdateToken(ctx contextx.Contextx, updated *user.Profile) (info *user.Profile, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := `update users set access_token=:access_token where id = :id`
	_, err = i.rw.NamedExecContext(timeout, stmt, updated)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
