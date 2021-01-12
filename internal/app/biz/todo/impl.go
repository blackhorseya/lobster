package todo

import (
	"github.com/blackhorseya/lobster/internal/app/biz/todo/repo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/todo"
)

type impl struct {
	repo repo.IRepo
}

// NewImpl serve caller to create an IBiz
func NewImpl(repo repo.IRepo) IBiz {
	return &impl{repo: repo}
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (*todo.Task, error) {
	// todo: 2021-01-13|07:36|doggy|implement me
	panic("implement me")
}

func (i *impl) List(ctx contextx.Contextx, page, size int) ([]*todo.Task, error) {
	// todo: 2021-01-13|07:36|doggy|implement me
	panic("implement me")
}

func (i *impl) Count(ctx contextx.Contextx) (int, error) {
	// todo: 2021-01-13|07:36|doggy|implement me
	panic("implement me")
}

func (i *impl) Create(ctx contextx.Contextx, task *todo.Task) (*todo.Task, error) {
	// todo: 2021-01-13|07:36|doggy|implement me
	panic("implement me")
}

func (i *impl) ChangeTitle(ctx contextx.Contextx, title string) (*todo.Task, error) {
	// todo: 2021-01-13|07:36|doggy|implement me
	panic("implement me")
}

func (i *impl) UpdateStatus(ctx contextx.Contextx, completed bool) (*todo.Task, error) {
	// todo: 2021-01-13|07:36|doggy|implement me
	panic("implement me")
}

func (i *impl) Delete(ctx contextx.Contextx, id string) error {
	// todo: 2021-01-13|07:36|doggy|implement me
	panic("implement me")
}
