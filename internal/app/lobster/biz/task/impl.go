package task

import (
	"time"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/task/repo"
	"github.com/blackhorseya/lobster/internal/pkg/entities/errors"
	"github.com/blackhorseya/lobster/internal/pkg/entities/task"
	"github.com/blackhorseya/lobster/internal/pkg/utils/contextx"
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
		i.logger.Error(errors.ErrInvalidID.Error(), zap.Error(err), zap.String("id", id))
		return nil, errors.ErrInvalidID
	}

	ret, err := i.repo.QueryByID(ctx, id)
	if err != nil {
		i.logger.Error(errors.ErrGetTaskByID.Error(), zap.Error(err), zap.String("id", id))
		return nil, errors.ErrGetTaskByID
	}
	if ret == nil {
		i.logger.Error(errors.ErrTaskNotExists.Error(), zap.Error(err), zap.String("id", id))
		return nil, errors.ErrTaskNotExists
	}

	return ret, nil
}

func (i *impl) List(ctx contextx.Contextx, page, size int) ([]*task.Task, error) {
	if page <= 0 {
		i.logger.Error(errors.ErrInvalidPage.Error(), zap.Int("page", page))
		return nil, errors.ErrInvalidPage
	}

	if size <= 0 {
		i.logger.Error(errors.ErrInvalidSize.Error(), zap.Int("size", size))
		return nil, errors.ErrInvalidSize
	}

	ret, err := i.repo.List(ctx, (page-1)*size, size)
	if err != nil {
		i.logger.Error(errors.ErrListTasks.Error(), zap.Error(err))
		return nil, errors.ErrTaskNotExists
	}

	return ret, nil
}

func (i *impl) Count(ctx contextx.Contextx) (int, error) {
	ret, err := i.repo.Count(ctx)
	if err != nil {
		i.logger.Error(errors.ErrCountObj.Error(), zap.Error(err))
		return 0, errors.ErrTaskNotExists
	}
	if ret == 0 {
		i.logger.Error("count all tasks is not found")
		return 0, errors.ErrTaskNotExists
	}

	return ret, nil
}

func (i *impl) Create(ctx contextx.Contextx, task *task.Task) (*task.Task, error) {
	if len(task.Title) == 0 {
		i.logger.Error(errors.ErrEmptyTitle.Error(), zap.String("title", task.Title))
		return nil, errors.ErrEmptyTitle
	}

	task.ID = uuid.New().String()
	task.CreateAt = time.Now().UnixNano()

	ret, err := i.repo.Create(ctx, task)
	if err != nil {
		i.logger.Error(errors.ErrCreateTask.Error(), zap.Error(err), zap.Any("created", task))
		return nil, errors.ErrCreateTask
	}

	return ret, nil
}

func (i *impl) UpdateStatus(ctx contextx.Contextx, id string, status task.Status) (t *task.Task, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		i.logger.Error(errors.ErrInvalidID.Error(), zap.Error(err), zap.String("id", id))
		return nil, err
	}

	exist, err := i.repo.QueryByID(ctx, id)
	if err != nil {
		i.logger.Error(errors.ErrGetTaskByID.Error(), zap.Error(err), zap.String("id", id))
		return nil, err
	}
	if exist == nil {
		i.logger.Error(errors.ErrTaskNotExists.Error(), zap.Error(err), zap.String("id", id))
		return nil, errors.ErrTaskNotExists
	}

	exist.Status = status
	ret, err := i.repo.Update(ctx, exist)
	if err != nil {
		i.logger.Error(errors.ErrUpdateTask.Error(), zap.Error(err), zap.String("id", id))
		return nil, err
	}

	return ret, nil
}

func (i *impl) ModifyTitle(ctx contextx.Contextx, id, title string) (t *task.Task, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		i.logger.Error(errors.ErrInvalidID.Error(), zap.Error(err), zap.String("id", id), zap.String("title", title))
		return nil, err
	}

	if len(title) == 0 {
		i.logger.Error(errors.ErrEmptyTitle.Error(), zap.String("id", id), zap.String("title", title))
		return nil, errors.ErrEmptyTitle
	}

	exist, err := i.repo.QueryByID(ctx, id)
	if err != nil {
		i.logger.Error(errors.ErrGetTaskByID.Error(), zap.Error(err), zap.String("id", id), zap.String("title", title))
		return nil, err
	}
	if exist == nil {
		i.logger.Error(errors.ErrTaskNotExists.Error(), zap.Error(err), zap.String("id", id), zap.String("title", title))
		return nil, errors.ErrTaskNotExists
	}

	exist.Title = title
	ret, err := i.repo.Update(ctx, exist)
	if err != nil {
		i.logger.Error(errors.ErrUpdateTask.Error(), zap.Error(err), zap.String("id", id), zap.String("title", title))
		return nil, err
	}

	return ret, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id string) error {
	if _, err := uuid.Parse(id); err != nil {
		i.logger.Error(errors.ErrInvalidID.Error(), zap.Error(err), zap.String("id", id))
		return errors.ErrInvalidID
	}

	ret, err := i.repo.Delete(ctx, id)
	if err != nil {
		i.logger.Error(errors.ErrTaskNotExists.Error(), zap.Error(err), zap.String("id", id))
		return errors.ErrTaskNotExists
	}
	if ret == 0 {
		i.logger.Error(errors.ErrTaskNotExists.Error(), zap.Error(err), zap.String("id", id))
		return errors.ErrTaskNotExists
	}

	return nil
}
