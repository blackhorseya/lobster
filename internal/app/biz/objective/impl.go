package objective

import (
	"github.com/blackhorseya/lobster/internal/app/biz/objective/repo"
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

func (i *impl) Create(ctx contextx.Contextx, obj *okr.Objective) (*okr.Objective, error) {
	panic("implement me")
}

func (i *impl) List(ctx contextx.Contextx, page, size int) ([]*okr.Objective, error) {
	panic("implement me")
}

func (i *impl) Count(ctx contextx.Contextx) (int, error) {
	panic("implement me")
}

func (i *impl) Update(ctx contextx.Contextx, updated *okr.Objective) (*okr.Objective, error) {
	panic("implement me")
}

func (i *impl) Delete(ctx contextx.Contextx, id string) error {
	panic("implement me")
}
