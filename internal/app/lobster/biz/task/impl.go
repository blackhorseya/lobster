package task

import (
	"time"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/task/repo"
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/er"
	"github.com/blackhorseya/lobster/internal/pkg/entity/task"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	repo   repo.IRepo
}

// NewImpl serve caller to create an IBiz
func NewImpl(logger *zap.Logger, repo repo.IRepo) IBiz {
	return &impl{
		logger: logger.With(zap.String("type", "TaskBiz")),
		repo:   repo,
	}
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (*task.Task, error) {
	if _, err := uuid.Parse(id); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err), zap.String("id", id))
		return nil, er.ErrInvalidID
	}

	ret, err := i.repo.QueryByID(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrGetTaskByID.Error(), zap.Error(err), zap.String("id", id))
		return nil, er.ErrGetTaskByID
	}
	if ret == nil {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.Error(err), zap.String("id", id))
		return nil, er.ErrTaskNotExists
	}

	return ret, nil
}

func (i *impl) List(ctx contextx.Contextx, page, size int) ([]*task.Task, error) {
	if page <= 0 {
		i.logger.Error(er.ErrInvalidPage.Error(), zap.Int("page", page))
		return nil, er.ErrInvalidPage
	}

	if size <= 0 {
		i.logger.Error(er.ErrInvalidSize.Error(), zap.Int("size", size))
		return nil, er.ErrInvalidSize
	}

	ret, err := i.repo.List(ctx, (page-1)*size, size)
	if err != nil {
		i.logger.Error(er.ErrListTasks.Error(), zap.Error(err))
		return nil, er.ErrTaskNotExists
	}

	return ret, nil
}

func (i *impl) Count(ctx contextx.Contextx) (int, error) {
	ret, err := i.repo.Count(ctx)
	if err != nil {
		i.logger.Error(er.ErrCountObj.Error(), zap.Error(err))
		return 0, er.ErrTaskNotExists
	}
	if ret == 0 {
		i.logger.Error("count all tasks is not found")
		return 0, er.ErrTaskNotExists
	}

	return ret, nil
}

func (i *impl) Create(ctx contextx.Contextx, task *task.Task) (*task.Task, error) {
	if len(task.Title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error(), zap.String("title", task.Title))
		return nil, er.ErrEmptyTitle
	}

	task.ID = uuid.New().String()
	task.CreatedAt = time.Now().UnixNano()

	ret, err := i.repo.Create(ctx, task)
	if err != nil {
		i.logger.Error(er.ErrCreateTask.Error(), zap.Error(err), zap.Any("created", task))
		return nil, er.ErrCreateTask
	}

	return ret, nil
}

func (i *impl) UpdateStatus(ctx contextx.Contextx, id string, status task.Status) (t *task.Task, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err), zap.String("id", id))
		return nil, err
	}

	exist, err := i.repo.QueryByID(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrGetTaskByID.Error(), zap.Error(err), zap.String("id", id))
		return nil, err
	}
	if exist == nil {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.Error(err), zap.String("id", id))
		return nil, er.ErrTaskNotExists
	}

	exist.Status = status
	ret, err := i.repo.Update(ctx, exist)
	if err != nil {
		i.logger.Error(er.ErrUpdateTask.Error(), zap.Error(err), zap.String("id", id))
		return nil, err
	}

	return ret, nil
}

func (i *impl) ModifyTitle(ctx contextx.Contextx, id, title string) (t *task.Task, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err), zap.String("id", id), zap.String("title", title))
		return nil, err
	}

	if len(title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error(), zap.String("id", id), zap.String("title", title))
		return nil, er.ErrEmptyTitle
	}

	exist, err := i.repo.QueryByID(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrGetTaskByID.Error(), zap.Error(err), zap.String("id", id), zap.String("title", title))
		return nil, err
	}
	if exist == nil {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.Error(err), zap.String("id", id), zap.String("title", title))
		return nil, er.ErrTaskNotExists
	}

	exist.Title = title
	ret, err := i.repo.Update(ctx, exist)
	if err != nil {
		i.logger.Error(er.ErrUpdateTask.Error(), zap.Error(err), zap.String("id", id), zap.String("title", title))
		return nil, err
	}

	return ret, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id string) error {
	if _, err := uuid.Parse(id); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err), zap.String("id", id))
		return er.ErrInvalidID
	}

	ret, err := i.repo.Delete(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.Error(err), zap.String("id", id))
		return er.ErrTaskNotExists
	}
	if ret == 0 {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.Error(err), zap.String("id", id))
		return er.ErrTaskNotExists
	}

	return nil
}
