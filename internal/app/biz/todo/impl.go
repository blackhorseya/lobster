package todo

import (
	"github.com/blackhorseya/lobster/internal/app/biz/todo/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/todo"
	er "github.com/blackhorseya/lobster/internal/pkg/entities/error"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type impl struct {
	repo repo.IRepo
}

// NewImpl serve caller to create an IBiz
func NewImpl(repo repo.IRepo) IBiz {
	return &impl{repo: repo}
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (*todo.Task, error) {
	if _, err := uuid.Parse(id); err != nil {
		ctx.WithFields(logrus.Fields{"err": err, "id": id}).Errorf("parse id is failure")
		return nil, er.ErrInvalidID
	}

	ret, err := i.repo.QueryByID(ctx, id)
	if err != nil {
		ctx.WithField("err", err).Errorf("query task by id is failure")
		return nil, er.ErrTaskNotExists
	}
	if ret == nil {
		ctx.WithField("id", id).Errorf("query task by id return empty")
		return nil, er.ErrTaskNotExists
	}

	return ret, nil
}

func (i *impl) List(ctx contextx.Contextx, page, size int) ([]*todo.Task, error) {
	if page <= 0 {
		ctx.WithField("page", page).Errorf("page is invalid")
		return nil, er.ErrInvalidPage
	}

	if size <= 0 {
		ctx.WithField("size", size).Errorf("size is invalid")
		return nil, er.ErrInvalidSize
	}

	ret, err := i.repo.List(ctx, (page-1)*size, size)
	if err != nil {
		ctx.WithField("err", err).Errorf("list all tasks is failure")
		return nil, er.ErrTaskNotExists
	}
	if len(ret) == 0 {
		ctx.WithFields(logrus.Fields{
			"page": page,
			"size": size,
		}).Errorf("list all tasks is not found")
		return nil, er.ErrTaskNotExists
	}

	return ret, nil
}

func (i *impl) Count(ctx contextx.Contextx) (int, error) {
	// todo: 2021-01-13|07:36|doggy|implement me
	panic("implement me")
}

func (i *impl) Create(ctx contextx.Contextx, task *todo.Task) (*todo.Task, error) {
	// todo: 2021-01-13|07:36|doggy|implement me
	panic("implement me")
}

func (i *impl) ChangeTitle(ctx contextx.Contextx, title string) (*todo.Task, error) {
	// todo: 2021-01-13|07:36|doggy|implement me
	panic("implement me")
}

func (i *impl) UpdateStatus(ctx contextx.Contextx, completed bool) (*todo.Task, error) {
	// todo: 2021-01-13|07:36|doggy|implement me
	panic("implement me")
}

func (i *impl) Delete(ctx contextx.Contextx, id string) error {
	// todo: 2021-01-13|07:36|doggy|implement me
	panic("implement me")
}
