package repo

import (
	"database/sql"
	"time"

	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/todo"
	"github.com/jmoiron/sqlx"
)

type impl struct {
	rw *sqlx.DB
}

// NewImpl serve caller to create an IRepo
func NewImpl(rw *sqlx.DB) IRepo {
	return &impl{rw: rw}
}

func (i *impl) QueryByID(ctx contextx.Contextx, userID, id int64) (*todo.Task, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	ret := todo.Task{}
	cmd := "SELECT id, result_id, title, status, created_at FROM tasks WHERE id = ? and user_id = ?"
	err := i.rw.GetContext(timeout, &ret, cmd, id, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &ret, nil
}

func (i *impl) Create(ctx contextx.Contextx, task *todo.Task) (*todo.Task, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cmd := "INSERT INTO tasks (id, result_id, title, status, created_at) VALUES (:id, :result_id, :title, :status, :created_at)"
	_, err := i.rw.NamedExecContext(timeout, cmd, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (i *impl) List(ctx contextx.Contextx, offset, limit int) ([]*todo.Task, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var ret []*todo.Task
	cmd := "SELECT id, result_id, title, status, created_at FROM tasks LIMIT ? OFFSET ?"
	if err := i.rw.SelectContext(timeout, &ret, cmd, limit, offset); err != nil {
		return nil, err
	}

	return ret, nil
}

func (i *impl) Count(ctx contextx.Contextx) (int, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var ret int
	cmd := "SELECT COUNT(*) FROM tasks"
	row := i.rw.QueryRowxContext(timeout, cmd)
	if err := row.Scan(&ret); err != nil {
		return 0, err
	}

	return ret, nil
}

func (i *impl) Update(ctx contextx.Contextx, updated *todo.Task) (*todo.Task, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cmd := "UPDATE tasks SET title=:title, status=:status WHERE id = :id"
	_, err := i.rw.NamedExecContext(timeout, cmd, updated)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id int64) (int, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cmd := "DELETE FROM tasks WHERE id = ?"
	_ = i.rw.QueryRowxContext(timeout, cmd, id)

	return 1, nil
}
