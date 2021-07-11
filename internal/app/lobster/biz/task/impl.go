package task

import (
	"time"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/task/repo"
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/er"
	"github.com/blackhorseya/lobster/internal/pkg/entity/todo"
	"github.com/blackhorseya/lobster/internal/pkg/entity/user"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	repo   repo.IRepo
	node   *snowflake.Node
}

// NewImpl serve caller to create an IBiz
func NewImpl(logger *zap.Logger, repo repo.IRepo, node *snowflake.Node) IBiz {
	return &impl{
		logger: logger.With(zap.String("type", "TaskBiz")),
		repo:   repo,
		node:   node,
	}
}

func (i *impl) GetByID(ctx contextx.Contextx, id int64) (*todo.Task, error) {
	info, ok := ctx.Value("user").(*user.Profile)
	if !ok {
		i.logger.Error(er.ErrUserNotExists.Error())
		return nil, er.ErrUserNotExists
	}

	ret, err := i.repo.QueryByID(ctx, info.ID, id)
	if err != nil {
		i.logger.Error(er.ErrGetTaskByID.Error(), zap.Error(err), zap.Int64("id", id))
		return nil, er.ErrGetTaskByID
	}
	if ret == nil {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.Error(err), zap.Int64("id", id))
		return nil, er.ErrTaskNotExists
	}

	return ret, nil
}

func (i *impl) List(ctx contextx.Contextx, page, size int) ([]*todo.Task, error) {
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

func (i *impl) Create(ctx contextx.Contextx, title string) (*todo.Task, error) {
	profile, ok := ctx.Value("user").(*user.Profile)
	if !ok {
		i.logger.Error(er.ErrUserNotExists.Error())
		return nil, er.ErrUserNotExists
	}

	if len(title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error(), zap.String("title", title))
		return nil, er.ErrEmptyTitle
	}

	created := &todo.Task{
		ID:        i.node.Generate().Int64(),
		UserID:    profile.ID,
		Title:     title,
		Status:    0,
		CreatedAt: time.Now().UnixNano(),
	}

	ret, err := i.repo.Create(ctx, created)
	if err != nil {
		i.logger.Error(er.ErrCreateTask.Error(), zap.Error(err), zap.Any("created", created))
		return nil, er.ErrCreateTask
	}

	return ret, nil
}

func (i *impl) UpdateStatus(ctx contextx.Contextx, id int64, status todo.Status) (t *todo.Task, err error) {
	info, ok := ctx.Value("user").(*user.Profile)
	if !ok {
		i.logger.Error(er.ErrUserNotExists.Error())
		return nil, er.ErrUserNotExists
	}

	exist, err := i.repo.QueryByID(ctx, info.ID, id)
	if err != nil {
		i.logger.Error(er.ErrGetTaskByID.Error(), zap.Error(err), zap.Int64("id", id))
		return nil, err
	}
	if exist == nil {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.Error(err), zap.Int64("id", id))
		return nil, er.ErrTaskNotExists
	}

	exist.Status = status
	ret, err := i.repo.Update(ctx, exist)
	if err != nil {
		i.logger.Error(er.ErrUpdateTask.Error(), zap.Error(err), zap.Int64("id", id))
		return nil, err
	}

	return ret, nil
}

func (i *impl) ModifyTitle(ctx contextx.Contextx, id int64, title string) (t *todo.Task, err error) {
	info, ok := ctx.Value("user").(*user.Profile)
	if !ok {
		i.logger.Error(er.ErrUserNotExists.Error())
		return nil, er.ErrUserNotExists
	}

	if len(title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error(), zap.Int64("id", id), zap.String("title", title))
		return nil, er.ErrEmptyTitle
	}

	exist, err := i.repo.QueryByID(ctx, info.ID, id)
	if err != nil {
		i.logger.Error(er.ErrGetTaskByID.Error(), zap.Error(err), zap.Int64("id", id), zap.String("title", title))
		return nil, err
	}
	if exist == nil {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.Error(err), zap.Int64("id", id), zap.String("title", title))
		return nil, er.ErrTaskNotExists
	}

	exist.Title = title
	ret, err := i.repo.Update(ctx, exist)
	if err != nil {
		i.logger.Error(er.ErrUpdateTask.Error(), zap.Error(err), zap.Int64("id", id), zap.String("title", title))
		return nil, err
	}

	return ret, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id int64) error {
	ret, err := i.repo.Delete(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.Error(err), zap.Int64("id", id))
		return er.ErrTaskNotExists
	}
	if ret == 0 {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.Error(err), zap.Int64("id", id))
		return er.ErrTaskNotExists
	}

	return nil
}
