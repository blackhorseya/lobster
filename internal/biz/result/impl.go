package result

import (
	"time"

	"github.com/blackhorseya/lobster/internal/biz/result/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
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

func (i *impl) List(ctx contextx.Contextx, page, size int) (krs []*okr.KeyResult, err error) {
	if page <= 0 {
		ctx.WithField("page", page).Error(er.ErrInvalidPage)
		return nil, er.ErrInvalidPage
	}

	if size <= 0 {
		ctx.WithField("size", size).Error(er.ErrInvalidSize)
		return nil, er.ErrInvalidSize
	}

	ret, err := i.repo.QueryList(ctx, (page-1)*size, size)
	if err != nil {
		ctx.WithError(err).WithFields(logrus.Fields{"page": page, "size": size}).Error(er.ErrListKeyResult)
		return nil, er.ErrListKeyResult
	}

	return ret, nil
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (kr *okr.KeyResult, err error) {
	if _, err = uuid.Parse(id); err != nil {
		ctx.WithError(err).WithField("id", id).Error(er.ErrInvalidID)
		return nil, er.ErrInvalidID
	}

	kr, err = i.repo.QueryByID(ctx, id)
	if err != nil {
		ctx.WithError(err).WithField("id", id).Error(er.ErrGetKRByID)
		return nil, er.ErrGetKRByID
	}
	if kr == nil {
		ctx.WithField("id", id).Error(er.ErrKRNotExists)
		return nil, er.ErrKRNotExists
	}

	return kr, nil
}

func (i *impl) GetByGoalID(ctx contextx.Contextx, id string) (krs []*okr.KeyResult, err error) {
	logger := ctx.WithField("id", id)

	_, err = uuid.Parse(id)
	if err != nil {
		logger.WithError(err).Error(er.ErrInvalidID)
		return nil, err
	}

	ret, err := i.repo.QueryByGoalID(ctx, id)
	if err != nil {
		logger.WithError(err).Error(er.ErrListKeyResult)
		return nil, err
	}

	return ret, nil
}

func (i *impl) LinkToGoal(ctx contextx.Contextx, created *okr.KeyResult) (kr *okr.KeyResult, err error) {
	logger := ctx.WithField("created", created)

	_, err = uuid.Parse(created.GoalID)
	if err != nil {
		logger.WithError(err).Error(er.ErrInvalidID)
		return nil, er.ErrInvalidID
	}

	if len(created.Title) == 0 {
		logger.Error(er.ErrEmptyTitle)
		return nil, er.ErrEmptyTitle
	}

	created.ID = uuid.New().String()
	created.CreateAt = time.Now().UnixNano()
	ret, err := i.repo.Create(ctx, created)
	if err != nil {
		logger.WithError(err).Error(er.ErrCreateKR)
		return nil, er.ErrCreateKR
	}

	return ret, nil
}

func (i *impl) Update(ctx contextx.Contextx, updated *okr.KeyResult) (kr *okr.KeyResult, err error) {
	logger := ctx.WithField("updated", updated)

	_, err = uuid.Parse(updated.ID)
	if err != nil {
		logger.WithError(err).Error(er.ErrInvalidID)
		return nil, er.ErrInvalidID
	}

	_, err = uuid.Parse(updated.GoalID)
	if err != nil {
		logger.WithError(err).Error(er.ErrInvalidID)
		return nil, er.ErrInvalidID
	}

	if len(updated.Title) == 0 {
		logger.Error(er.ErrEmptyTitle)
		return nil, er.ErrEmptyTitle
	}

	exist, err := i.repo.QueryByID(ctx, updated.ID)
	if err != nil {
		logger.WithError(err).Error(er.ErrGetKRByID)
		return nil, er.ErrGetKRByID
	}
	if exist == nil {
		logger.Error(er.ErrKRNotExists)
		return nil, er.ErrKRNotExists
	}

	updated.ID = exist.ID
	updated.CreateAt = exist.CreateAt
	ret, err := i.repo.Update(ctx, updated)
	if err != nil {
		logger.WithError(err).Error(er.ErrUpdateKeyResult)
		return nil, er.ErrUpdateKeyResult
	}

	return ret, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id string) (err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		ctx.WithError(err).WithField("id", id).Error(er.ErrInvalidID)
		return er.ErrInvalidID
	}

	err = i.repo.Delete(ctx, id)
	if err != nil {
		ctx.WithError(err).WithField("id", id).Error(er.ErrDeleteKeyResult)
		return er.ErrDeleteKeyResult
	}

	return nil
}
