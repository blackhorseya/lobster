package objective

import (
	"time"

	"github.com/blackhorseya/lobster/internal/app/biz/objective/repo"
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

func (i *impl) Create(ctx contextx.Contextx, obj *okr.Objective) (*okr.Objective, error) {
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

func (i *impl) List(ctx contextx.Contextx, page, size int) ([]*okr.Objective, error) {
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
	if len(ret) == 0 {
		ctx.WithFields(logrus.Fields{"page": page, "size": size}).Error(er.ErrObjectiveNotExists)
		return nil, er.ErrObjectiveNotExists
	}

	return ret, nil
}

func (i *impl) Count(ctx contextx.Contextx) (int, error) {
	// todo: 2021-01-19|11:20|doggy|implement me
	panic("implement me")
}

func (i *impl) Update(ctx contextx.Contextx, updated *okr.Objective) (*okr.Objective, error) {
	// todo: 2021-01-19|11:20|doggy|implement me
	panic("implement me")
}

func (i *impl) Delete(ctx contextx.Contextx, id string) error {
	// todo: 2021-01-19|11:20|doggy|implement me
	panic("implement me")
}
