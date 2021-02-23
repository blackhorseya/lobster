package goal

import (
	"time"

	"github.com/blackhorseya/lobster/internal/biz/goal/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	er "github.com/blackhorseya/lobster/internal/pkg/entities/error"
	"github.com/blackhorseya/lobster/internal/pkg/pb"
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

func (i *impl) Create(ctx contextx.Contextx, obj *pb.Objective) (*pb.Objective, error) {
	if len(obj.Title) == 0 {
		ctx.WithField("title", obj.Title).Error(er.ErrEmptyTitle)
		return nil, er.ErrEmptyTitle
	}

	obj.ID = uuid.New().String()
	obj.CreateAt = time.Now().UnixNano()

	ret, err := i.repo.Create(ctx, obj)
	if err != nil {
		ctx.WithField("err", err).Error(er.ErrCreateObjective)
		return nil, er.ErrCreateObjective
	}

	return ret, nil
}

func (i *impl) List(ctx contextx.Contextx, page, size int) ([]*pb.Objective, error) {
	if page <= 0 {
		ctx.WithField("page", page).Error(er.ErrInvalidPage)
		return nil, er.ErrInvalidPage
	}

	if size <= 0 {
		ctx.WithField("size", size).Error(er.ErrInvalidSize)
		return nil, er.ErrInvalidSize
	}

	ret, err := i.repo.List(ctx, (page-1)*size, size)
	if err != nil {
		ctx.WithField("err", err).Error(er.ErrListObjectives)
		return nil, er.ErrListObjectives
	}

	return ret, nil
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (*pb.Objective, error) {
	if _, err := uuid.Parse(id); err != nil {
		ctx.WithFields(logrus.Fields{"err": err, "id": id}).Error(er.ErrInvalidID)
		return nil, er.ErrInvalidID
	}

	ret, err := i.repo.QueryByID(ctx, id)
	if err != nil {
		ctx.WithFields(logrus.Fields{"err": err, "id": id}).Error(er.ErrGetObjByID)
		return nil, er.ErrGetObjByID
	}

	return ret, nil
}

func (i *impl) Count(ctx contextx.Contextx) (int, error) {
	ret, err := i.repo.Count(ctx)
	if err != nil {
		ctx.WithField("err", err).Error(er.ErrCountObjective)
		return 0, er.ErrCountObjective
	}

	return ret, nil
}

func (i *impl) ModifyTitle(ctx contextx.Contextx, id, title string) (obj *pb.Objective, err error) {
	logger := ctx.WithField("id", id).WithField("title", title)

	_, err = uuid.Parse(id)
	if err != nil {
		logger.WithError(err).Error(er.ErrInvalidID)
		return nil, err
	}

	if len(title) == 0 {
		logger.Error(er.ErrEmptyTitle)
		return nil, er.ErrEmptyTitle
	}

	exist, err := i.repo.QueryByID(ctx, id)
	if err != nil {
		logger.WithError(err).Error(er.ErrGetObjByID)
		return nil, err
	}
	if exist == nil {
		logger.Error(er.ErrObjectiveNotExists)
		return nil, er.ErrObjectiveNotExists
	}

	exist.Title = title
	ret, err := i.repo.Update(ctx, exist)
	if err != nil {
		logger.WithError(err).Error(er.ErrUpdateObj)
		return nil, err
	}

	return ret, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id string) error {
	if _, err := uuid.Parse(id); err != nil {
		ctx.WithField("err", err).Error(er.ErrInvalidID)
		return er.ErrInvalidID
	}

	ret, err := i.repo.Delete(ctx, id)
	if err != nil {
		ctx.WithField("err", err).Error(er.ErrDeleteObj)
		return er.ErrDeleteObj
	}
	if ret == 0 {
		ctx.WithField("id", id).Error(er.ErrObjectiveNotExists)
		return er.ErrObjectiveNotExists
	}

	return nil
}
