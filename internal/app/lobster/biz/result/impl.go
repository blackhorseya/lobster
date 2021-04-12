package result

import (
	"time"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/result/repo"
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
		logger: logger.With(zap.String("type", "ResultBiz")),
		repo:   repo,
	}
}

func (i *impl) List(ctx contextx.Contextx, page, size int) (krs []*okr.Result, err error) {
	if page <= 0 {
		i.logger.Error(errors.ErrInvalidPage.Error(), zap.Int("page", page))
		return nil, errors.ErrInvalidPage
	}

	if size <= 0 {
		i.logger.Error(errors.ErrInvalidSize.Error(), zap.Int("size", size))
		return nil, errors.ErrInvalidSize
	}

	ret, err := i.repo.QueryList(ctx, (page-1)*size, size)
	if err != nil {
		i.logger.Error(errors.ErrListKR.Error(), zap.Error(err), zap.Int("page", page), zap.Int("size", size))
		return nil, errors.ErrListKR
	}

	return ret, nil
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (kr *okr.Result, err error) {
	if _, err = uuid.Parse(id); err != nil {
		i.logger.Error(errors.ErrInvalidID.Error(), zap.Error(err), zap.String("id", id))
		return nil, errors.ErrInvalidID
	}

	kr, err = i.repo.QueryByID(ctx, id)
	if err != nil {
		i.logger.Error(errors.ErrGetKRByID.Error(), zap.Error(err), zap.String("id", id))
		return nil, errors.ErrGetKRByID
	}
	if kr == nil {
		i.logger.Error(errors.ErrKRNotExists.Error(), zap.String("id", id))
		return nil, errors.ErrKRNotExists
	}

	return kr, nil
}

func (i *impl) GetByGoalID(ctx contextx.Contextx, id string) (krs []*okr.Result, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		i.logger.Error(errors.ErrInvalidID.Error(), zap.Error(err))
		return nil, err
	}

	ret, err := i.repo.QueryByGoalID(ctx, id)
	if err != nil {
		i.logger.Error(errors.ErrListKR.Error(), zap.Error(err))
		return nil, err
	}

	return ret, nil
}

func (i *impl) LinkToGoal(ctx contextx.Contextx, created *okr.Result) (kr *okr.Result, err error) {
	_, err = uuid.Parse(created.GoalID)
	if err != nil {
		i.logger.Error(errors.ErrInvalidID.Error(), zap.Error(err))
		return nil, errors.ErrInvalidID
	}

	if len(created.Title) == 0 {
		i.logger.Error(errors.ErrEmptyTitle.Error())
		return nil, errors.ErrEmptyTitle
	}

	created.ID = uuid.New().String()
	created.CreateAt = time.Now().UnixNano()
	ret, err := i.repo.Create(ctx, created)
	if err != nil {
		i.logger.Error(errors.ErrCreateKR.Error(), zap.Error(err))
		return nil, errors.ErrCreateKR
	}

	return ret, nil
}

func (i *impl) ModifyTitle(ctx contextx.Contextx, id, title string) (result *okr.Result, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		i.logger.Error(errors.ErrInvalidID.Error())
		return nil, errors.ErrInvalidID
	}

	if len(title) == 0 {
		i.logger.Error(errors.ErrEmptyTitle.Error())
		return nil, errors.ErrEmptyTitle
	}

	exist, err := i.repo.QueryByID(ctx, id)
	if err != nil {
		i.logger.Error(errors.ErrGetKRByID.Error(), zap.Error(err))
		return nil, errors.ErrGetKRByID
	}
	if exist == nil {
		i.logger.Error(errors.ErrKRNotExists.Error())
		return nil, errors.ErrKRNotExists
	}

	exist.Title = title
	ret, err := i.repo.Update(ctx, exist)
	if err != nil {
		i.logger.Error(errors.ErrUpdateKR.Error(), zap.Error(err))
		return nil, errors.ErrUpdateKR
	}

	return ret, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id string) (err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		i.logger.Error(errors.ErrInvalidID.Error(), zap.Error(err), zap.String("id", id))
		return errors.ErrInvalidID
	}

	err = i.repo.Delete(ctx, id)
	if err != nil {
		i.logger.Error(errors.ErrDeleteKR.Error(), zap.Error(err), zap.String("id", id))
		return errors.ErrDeleteKR
	}

	return nil
}
