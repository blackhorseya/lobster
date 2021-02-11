package kr

import (
	"github.com/blackhorseya/lobster/internal/biz/kr"
	"github.com/gin-gonic/gin"
)

type impl struct {
	biz kr.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(biz kr.IBiz) IHandler {
	return &impl{biz: biz}
}

func (i *impl) GetByID(c *gin.Context) {
	panic("implement me")
}

func (i *impl) List(c *gin.Context) {
	panic("implement me")
}

func (i *impl) Create(c *gin.Context) {
	panic("implement me")
}

func (i *impl) Update(c *gin.Context) {
	panic("implement me")
}

func (i *impl) Delete(c *gin.Context) {
	panic("implement me")
}
