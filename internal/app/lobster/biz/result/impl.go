package result

import (
	"time"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/result/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	er "github.com/blackhorseya/lobster/internal/pkg/entities/error"
	"github.com/blackhorseya/lobster/internal/pkg/entities/okr"
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
		i.logger.Error(er.ErrInvalidPage.Error(), zap.Int("page", page))
		return nil, er.ErrInvalidPage
	}

	if size <= 0 {
		i.logger.Error(er.ErrInvalidSize.Error(), zap.Int("size", size))
		return nil, er.ErrInvalidSize
	}

	ret, err := i.repo.QueryList(ctx, (page-1)*size, size)
	if err != nil {
		i.logger.Error(er.ErrListKeyResult.Error(), zap.Error(err), zap.Int("page", page), zap.Int("size", size))
		return nil, er.ErrListKeyResult
	}

	return ret, nil
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (kr *okr.Result, err error) {
	if _, err = uuid.Parse(id); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err), zap.String("id", id))
		return nil, er.ErrInvalidID
	}

	kr, err = i.repo.QueryByID(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrGetKRByID.Error(), zap.Error(err), zap.String("id", id))
		return nil, er.ErrGetKRByID
	}
	if kr == nil {
		i.logger.Error(er.ErrKRNotExists.Error(), zap.String("id", id))
		return nil, er.ErrKRNotExists
	}

	return kr, nil
}

func (i *impl) GetByGoalID(ctx contextx.Contextx, id string) (krs []*okr.Result, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		return nil, err
	}

	ret, err := i.repo.QueryByGoalID(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrListKeyResult.Error(), zap.Error(err))
		return nil, err
	}

	return ret, nil
}

func (i *impl) LinkToGoal(ctx contextx.Contextx, created *okr.Result) (kr *okr.Result, err error) {
	_, err = uuid.Parse(created.GoalID)
	if err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		return nil, er.ErrInvalidID
	}

	if len(created.Title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error())
		return nil, er.ErrEmptyTitle
	}

	created.ID = uuid.New().String()
	created.CreateAt = time.Now().UnixNano()
	ret, err := i.repo.Create(ctx, created)
	if err != nil {
		i.logger.Error(er.ErrCreateKR.Error(), zap.Error(err))
		return nil, er.ErrCreateKR
	}

	return ret, nil
}

func (i *impl) ModifyTitle(ctx contextx.Contextx, id, title string) (result *okr.Result, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		i.logger.Error(er.ErrInvalidID.Error())
		return nil, er.ErrInvalidID
	}

	if len(title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error())
		return nil, er.ErrEmptyTitle
	}

	exist, err := i.repo.QueryByID(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrGetKRByID.Error(), zap.Error(err))
		return nil, er.ErrGetKRByID
	}
	if exist == nil {
		i.logger.Error(er.ErrKRNotExists.Error())
		return nil, er.ErrKRNotExists
	}

	exist.Title = title
	ret, err := i.repo.Update(ctx, exist)
	if err != nil {
		i.logger.Error(er.ErrUpdateKeyResult.Error(), zap.Error(err))
		return nil, er.ErrUpdateKeyResult
	}

	return ret, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id string) (err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err), zap.String("id", id))
		return er.ErrInvalidID
	}

	err = i.repo.Delete(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrDeleteKeyResult.Error(), zap.Error(err), zap.String("id", id))
		return er.ErrDeleteKeyResult
	}

	return nil
}
