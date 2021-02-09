package repo

import (
	"time"

	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"github.com/jmoiron/sqlx"
)

type impl struct {
	rw *sqlx.DB
}

// NewImpl serve caller to create an IRepo
func NewImpl(rw *sqlx.DB) IRepo {
	return &impl{rw: rw}
}

func (i *impl) QueryKRByID(ctx contextx.Contextx, goalID, krID string) (*okr.KeyResult, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var kr okr.KeyResult
	stmt := `select id, title, target, actual, create_at from keyresults where id = ? and goal_id = ?`
	err := i.rw.GetContext(timeout, &kr, stmt, krID, goalID)
	if err != nil {
		ctx.WithError(err).Errorln("query key result by id is failure")
		return nil, err
	}

	return &kr, nil
}
