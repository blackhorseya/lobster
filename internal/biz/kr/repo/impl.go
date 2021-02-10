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

func (i *impl) QueryByID(ctx contextx.Contextx, id string) (*okr.KeyResult, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var kr okr.KeyResult
	stmt := `select id, title, target, actual, create_at from keyresults where id = ?`
	err := i.rw.GetContext(timeout, &kr, stmt, id)
	if err != nil {
		ctx.WithError(err).Errorln("query key result by id is failure")
		return nil, err
	}

	return &kr, nil
}

func (i *impl) Create(ctx contextx.Contextx, created *okr.KeyResult) (kr *okr.KeyResult, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := `insert into keyresults (id, goal_id, title, target, actual, create_at)
values (:id, :goal_id, :title, :target, :actual, :create_at)`
	_, err = i.rw.NamedExecContext(timeout, stmt, created)
	if err != nil {
		ctx.WithError(err).Errorln("insert key result info is failure")
		return nil, err
	}

	return created, nil
}

func (i *impl) QueryList(ctx contextx.Contextx, offset, limit int) (krs []*okr.KeyResult, err error) {
	// todo: 2021-02-10|09:28|doggy|implement me
	panic("implement me")
}

func (i *impl) Delete(ctx contextx.Contextx, id string) (err error) {
	// todo: 2021-02-10|09:28|doggy|implement me
	panic("implement me")
}
