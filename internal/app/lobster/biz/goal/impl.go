package goal

import (
	"time"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/goal/repo"
	"github.com/blackhorseya/lobster/internal/pkg/entities/errors"
	"github.com/blackhorseya/lobster/internal/pkg/entities/okr"
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
		logger: logger.With(zap.String("type", "GoalBiz")),
		repo:   repo,
	}
}

func (i *impl) Create(ctx contextx.Contextx, obj *okr.Goal) (*okr.Goal, error) {
	if len(obj.Title) == 0 {
		i.logger.Error(errors.ErrEmptyTitle.Error(), zap.String("title", obj.Title))
		return nil, errors.ErrEmptyTitle
	}

	obj.ID = uuid.New().String()
	obj.CreateAt = time.Now().UnixNano()

	ret, err := i.repo.Create(ctx, obj)
	if err != nil {
		i.logger.Error(errors.ErrCreateObj.Error())
		return nil, errors.ErrCreateObj
	}

	return ret, nil
}

func (i *impl) List(ctx contextx.Contextx, page, size int) ([]*okr.Goal, error) {
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
		i.logger.Error(errors.ErrListObj.Error(), zap.Error(err))
		return nil, errors.ErrListObj
	}

	return ret, nil
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (*okr.Goal, error) {
	if _, err := uuid.Parse(id); err != nil {
		i.logger.Error(errors.ErrInvalidID.Error(), zap.Error(err), zap.String("id", id))
		return nil, errors.ErrInvalidID
	}

	ret, err := i.repo.QueryByID(ctx, id)
	if err != nil {
		i.logger.Error(errors.ErrGetObjByID.Error(), zap.Error(err), zap.String("id", id))
		return nil, errors.ErrGetObjByID
	}

	return ret, nil
}

func (i *impl) Count(ctx contextx.Contextx) (int, error) {
	ret, err := i.repo.Count(ctx)
	if err != nil {
		i.logger.Error(errors.ErrCountObj.Error(), zap.Error(err))
		return 0, errors.ErrCountObj
	}

	return ret, nil
}

func (i *impl) ModifyTitle(ctx contextx.Contextx, id, title string) (obj *okr.Goal, err error) {
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
		i.logger.Error(errors.ErrGetObjByID.Error(), zap.Error(err), zap.String("id", id), zap.String("title", title))
		return nil, err
	}
	if exist == nil {
		i.logger.Error(errors.ErrObjNotExists.Error(), zap.String("id", id), zap.String("title", title))
		return nil, errors.ErrObjNotExists
	}

	exist.Title = title
	ret, err := i.repo.Update(ctx, exist)
	if err != nil {
		i.logger.Error(errors.ErrUpdateObj.Error(), zap.Error(err), zap.String("id", id), zap.String("title", title))
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
		i.logger.Error(errors.ErrDeleteObj.Error(), zap.Error(err), zap.String("id", id))
		return errors.ErrDeleteObj
	}
	if ret == 0 {
		i.logger.Error(errors.ErrObjNotExists.Error(), zap.String("id", id))
		return errors.ErrObjNotExists
	}

	return nil
}
