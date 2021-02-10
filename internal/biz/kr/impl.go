package kr

import (
	"github.com/blackhorseya/lobster/internal/biz/kr/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
)

type impl struct {
	repo repo.IRepo
}

// NewImpl serve caller to create an IBiz
func NewImpl(repo repo.IRepo) IBiz {
	return &impl{repo: repo}
}

func (i *impl) List(ctx contextx.Contextx, page, size int) (krs []*okr.KeyResult, err error) {
	panic("implement me")
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (kr *okr.KeyResult, err error) {
	panic("implement me")
}

func (i *impl) LinkToGoal(ctx contextx.Contextx, created *okr.KeyResult) (kr *okr.KeyResult, err error) {
	panic("implement me")
}

func (i *impl) Update(ctx contextx.Contextx, updated *okr.KeyResult) (kr *okr.KeyResult, err error) {
	panic("implement me")
}

func (i *impl) Delete(ctx contextx.Contextx, id string) (err error) {
	panic("implement me")
}
