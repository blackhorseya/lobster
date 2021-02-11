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
		ctx.WithError(err).WithField("created", created).Errorln("insert key result is failure")
		return nil, err
	}

	return created, nil
}

func (i *impl) Update(ctx contextx.Contextx, updated *okr.KeyResult) (kr *okr.KeyResult, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := `update keyresults set goal_id=:goal_id, title=:title, target=:target, actual=:actual where id = :id`
	_, err = i.rw.NamedExecContext(timeout, stmt, updated)
	if err != nil {
		ctx.WithError(err).WithField("updated", updated).Errorln("update key result is failure")
		return nil, err
	}

	return updated, nil
}

func (i *impl) QueryList(ctx contextx.Contextx, offset, limit int) (krs []*okr.KeyResult, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var ret []*okr.KeyResult
	stmt := `select id, goal_id, title, target, actual, create_at from keyresults limit ? offset ?`
	err = i.rw.SelectContext(timeout, &ret, stmt, limit, offset)
	if err != nil {
		ctx.WithError(err).Errorln("list all key results by limit and offset is failure")
		return nil, err
	}

	return ret, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id string) (err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := `delete from keyresults where id = ?`
	_ = i.rw.QueryRowxContext(timeout, stmt, id)

	return nil
}
