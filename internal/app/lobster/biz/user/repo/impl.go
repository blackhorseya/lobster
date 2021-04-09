package repo

import (
	"time"

	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/user"
	"github.com/jmoiron/sqlx"
)

type impl struct {
	rw *sqlx.DB
}

// NewImpl serve caller to create an IRepo
func NewImpl(rw *sqlx.DB) IRepo {
	return &impl{rw: rw}
}

func (i *impl) QueryInfoByEmail(ctx contextx.Contextx, email string) (info *user.Profile, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	ret := user.Profile{}
	stmt := `select sn, access_token, email, signup_at from users where email = ?`
	err = i.rw.GetContext(timeout, &ret, stmt, email)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (i *impl) UserRegister(ctx contextx.Contextx, newUser user.Profile) (info *user.Profile, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := `insert into users (sn, access_token, email, signup_at) values (:sn, :access_token, :email, :signup_at)`
	_, err = i.rw.NamedExecContext(timeout, stmt, newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
