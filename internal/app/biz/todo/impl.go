package todo

import (
	"time"

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
	ret, err := i.repo.Count(ctx)
	if err != nil {
		ctx.WithField("err", err).Errorf("count all tasks is failure")
		return 0, er.ErrTaskNotExists
	}
	if ret == 0 {
		ctx.Errorf("count all tasks is not found")
		return 0, er.ErrTaskNotExists
	}

	return ret, nil
}

func (i *impl) Create(ctx contextx.Contextx, task *todo.Task) (*todo.Task, error) {
	if len(task.Title) == 0 {
		ctx.WithField("title", task.Title).Errorf(er.ErrEmptyTitle.Error())
		return nil, er.ErrEmptyTitle
	}

	task.ID = uuid.New().String()
	task.CreateAt = time.Now().UnixNano()

	ret, err := i.repo.Create(ctx, task)
	if err != nil {
		ctx.WithFields(logrus.Fields{
			"err":     err,
			"created": task,
		}).Errorf(er.ErrCreateTask.Error())
		return nil, er.ErrCreateTask
	}

	return ret, nil
}

func (i *impl) ChangeTitle(ctx contextx.Contextx, id, title string) (*todo.Task, error) {
	if _, err := uuid.Parse(id); err != nil {
		ctx.WithFields(logrus.Fields{"err": err, "id": id}).Error(er.ErrInvalidID)
		return nil, er.ErrInvalidID
	}

	if len(title) == 0 {
		ctx.WithField("title", title).Error(er.ErrEmptyTitle)
		return nil, er.ErrEmptyTitle
	}

	exist, err := i.repo.QueryByID(ctx, id)
	if err != nil {
		ctx.WithFields(logrus.Fields{"err": err, "id": id}).Error(er.ErrTaskNotExists)
		return nil, er.ErrTaskNotExists
	}
	if exist == nil {
		ctx.WithField("id", id).Error(er.ErrTaskNotExists)
		return nil, er.ErrTaskNotExists
	}

	exist.Title = title

	ret, err := i.repo.Update(ctx, exist)
	if err != nil {
		ctx.WithField("err", err).Error(er.ErrUpdateTask)
		return nil, er.ErrUpdateTask
	}

	return ret, nil
}

func (i *impl) UpdateStatus(ctx contextx.Contextx, id string, completed bool) (*todo.Task, error) {
	if _, err := uuid.Parse(id); err != nil {
		ctx.WithFields(logrus.Fields{"err": err, "id": id}).Error(er.ErrInvalidID)
		return nil, er.ErrInvalidID
	}

	exist, err := i.repo.QueryByID(ctx, id)
	if err != nil {
		ctx.WithFields(logrus.Fields{"err": err, "id": id}).Error(er.ErrTaskNotExists)
		return nil, er.ErrTaskNotExists
	}
	if exist == nil {
		ctx.WithField("id", id).Error(er.ErrTaskNotExists)
		return nil, er.ErrTaskNotExists
	}

	exist.Completed = completed

	ret, err := i.repo.Update(ctx, exist)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id string) error {
	if _, err := uuid.Parse(id); err != nil {
		ctx.WithFields(logrus.Fields{"err": err, "id": id}).Error(er.ErrInvalidID)
		return er.ErrInvalidID
	}

	ret, err := i.repo.Delete(ctx, id)
	if err != nil {
		ctx.WithFields(logrus.Fields{"err": err, "id": id}).Error(er.ErrTaskNotExists)
		return er.ErrTaskNotExists
	}
	if ret == 0 {
		ctx.WithField("id", id).Error(er.ErrTaskNotExists)
		return er.ErrTaskNotExists
	}

	return nil
}
