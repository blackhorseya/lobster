package task

import (
	"time"

	"github.com/blackhorseya/lobster/internal/biz/task/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities"
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

func (i *impl) GetByID(ctx contextx.Contextx, id string) (*task.Task, error) {
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

func (i *impl) List(ctx contextx.Contextx, page, size int) ([]*task.Task, error) {
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

func (i *impl) Create(ctx contextx.Contextx, task *task.Task) (*task.Task, error) {
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

func (i *impl) Update(ctx contextx.Contextx, updated *task.Task) (*task.Task, error) {
	if _, err := uuid.Parse(updated.ID); err != nil {
		ctx.WithFields(logrus.Fields{"err": err, "id": updated.ID}).Error(er.ErrInvalidID)
		return nil, err
	}

	if len(updated.Title) == 0 {
		ctx.WithFields(logrus.Fields{"title": updated.Title}).Error(er.ErrEmptyTitle)
		return nil, er.ErrEmptyTitle
	}

	exist, err := i.repo.QueryByID(ctx, updated.ID)
	if err != nil {
		ctx.WithField("err", err).Error(er.ErrGetTaskByID)
		return nil, err
	}
	if exist == nil {
		ctx.WithField("updated", updated).Error(er.ErrTaskNotExists)
		return nil, er.ErrTaskNotExists
	}

	updated.CreateAt = exist.CreateAt
	ret, err := i.repo.Update(ctx, updated)
	if err != nil {
		ctx.WithFields(logrus.Fields{"err": err}).Error(er.ErrUpdateTask)
		return nil, er.ErrUpdateTask
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
