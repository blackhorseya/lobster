package okr

import (
	"github.com/blackhorseya/lobster/internal/app/biz/okr/repo"
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

// todo: 2021-01-18|21:47|doggy|implement me
func (i *impl) SetObjective(ctx contextx.Contextx, obj *okr.Objective) (*okr.Objective, error) {
	panic("implement me")
}

func (i *impl) ListObjectives(ctx contextx.Contextx, page, size int) ([]*okr.Objective, error) {
	panic("implement me")
}

func (i *impl) CountObjective(ctx contextx.Contextx) (int, error) {
	panic("implement me")
}

func (i *impl) UpdateObjective(ctx contextx.Contextx, updated *okr.Objective) (*okr.Objective, error) {
	panic("implement me")
}

func (i *impl) DeleteObjective(ctx contextx.Contextx, id string) error {
	panic("implement me")
}
