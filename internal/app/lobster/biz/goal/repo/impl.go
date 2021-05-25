package repo

import (
	"time"

	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/okr"
	"github.com/jmoiron/sqlx"
)

type impl struct {
	rw *sqlx.DB
}

// NewImpl serve caller to create an IRepo
func NewImpl(rw *sqlx.DB) IRepo {
	return &impl{rw: rw}
}

func (i *impl) Create(ctx contextx.Contextx, created *okr.Goal) (*okr.Goal, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cmd := `INSERT INTO goals (id, title, start_at, end_at, create_at) VALUES (:id, :title, :start_at, :end_at, :create_at)`
	_, err := i.rw.NamedExecContext(timeout, cmd, created)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (i *impl) QueryByID(ctx contextx.Contextx, id string) (*okr.Goal, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var ret okr.Goal
	cmd := "SELECT id, title, start_at, end_at, create_at FROM goals WHERE id = ?"
	err := i.rw.GetContext(timeout, &ret, cmd, id)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (i *impl) List(ctx contextx.Contextx, offset, limit int) ([]*okr.Goal, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var ret []*okr.Goal
	cmd := "SELECT id, title, start_at, end_at, create_at FROM goals LIMIT ? OFFSET ?"
	err := i.rw.SelectContext(timeout, &ret, cmd, limit, offset)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (i *impl) Count(ctx contextx.Contextx) (int, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var ret int
	cmd := "SELECT COUNT(*) FROM goals"
	row := i.rw.QueryRowxContext(timeout, cmd)
	err := row.Scan(&ret)
	if err != nil {
		return 0, err
	}

	return ret, nil
}

func (i *impl) Update(ctx contextx.Contextx, updated *okr.Goal) (*okr.Goal, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cmd := "UPDATE goals SET title=:title, start_at=:start_at, end_at=:end_at WHERE id = :id"
	_, err := i.rw.NamedExecContext(timeout, cmd, updated)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id string) (int, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cmd := "DELETE FROM goals WHERE id = ?"
	_ = i.rw.QueryRowxContext(timeout, cmd, id)

	return 1, nil
}
