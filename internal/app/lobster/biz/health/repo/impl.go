package repo

import (
	"time"

	"github.com/blackhorseya/lobster/internal/pkg/utils/contextx"
	"github.com/jmoiron/sqlx"
)

type impl struct {
	rw *sqlx.DB
}

// NewImpl serve caller to create an IRepo
func NewImpl(db *sqlx.DB) IRepo {
	return &impl{rw: db}
}

func (i *impl) Ping(ctx contextx.Contextx) (bool, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := i.rw.PingContext(timeout); err != nil {
		return false, err
	}

	return true, nil
}
