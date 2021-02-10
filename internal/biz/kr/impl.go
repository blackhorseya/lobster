package kr

import (
	"errors"

	"github.com/blackhorseya/lobster/internal/biz/kr/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"github.com/google/uuid"
)

type impl struct {
	repo repo.IRepo
}

// NewImpl serve caller to create an IBiz
func NewImpl(repo repo.IRepo) IBiz {
	return &impl{repo: repo}
}

func (i *impl) List(ctx contextx.Contextx, page, size int) (krs []*okr.KeyResult, err error) {
	// todo: 2021-02-10|09:28|doggy|implement me
	panic("implement me")
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (kr *okr.KeyResult, err error) {
	if _, err = uuid.Parse(id); err != nil {
		ctx.WithError(err).WithField("id", id).Error("parse id to uuid is failure")
		return nil, err
	}

	kr, err = i.repo.QueryByID(ctx, id)
	if err != nil {
		ctx.WithError(err).WithField("id", id).Error("query by id is failure")
		return nil, err
	}
	if kr == nil {
		ctx.WithField("id", id).Error("query by id is not found")
		return nil, errors.New("query by id is not found")
	}

	return kr, nil
}

func (i *impl) LinkToGoal(ctx contextx.Contextx, created *okr.KeyResult) (kr *okr.KeyResult, err error) {
	// todo: 2021-02-10|09:28|doggy|implement me
	panic("implement me")
}

func (i *impl) Update(ctx contextx.Contextx, updated *okr.KeyResult) (kr *okr.KeyResult, err error) {
	// todo: 2021-02-10|09:28|doggy|implement me
	panic("implement me")
}

func (i *impl) Delete(ctx contextx.Contextx, id string) (err error) {
	// todo: 2021-02-10|09:28|doggy|implement me
	panic("implement me")
}
