package todo

import (
	"github.com/blackhorseya/lobster/internal/app/biz/todo"
	"github.com/gin-gonic/gin"
)

type impl struct {
	biz todo.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(biz todo.IBiz) IHandler {
	return &impl{biz: biz}
}

func (i *impl) GetByID(c *gin.Context) {
	// todo: 2021-01-13|23:17|doggy|implement me
	panic("implement me")
}

func (i *impl) List(c *gin.Context) {
	// todo: 2021-01-13|23:17|doggy|implement me
	panic("implement me")
}

func (i *impl) Create(c *gin.Context) {
	// todo: 2021-01-13|23:17|doggy|implement me
	panic("implement me")
}

func (i *impl) Update(c *gin.Context) {
	// todo: 2021-01-13|23:17|doggy|implement me
	panic("implement me")
}

func (i *impl) Delete(c *gin.Context) {
	// todo: 2021-01-13|23:17|doggy|implement me
	panic("implement me")
}
