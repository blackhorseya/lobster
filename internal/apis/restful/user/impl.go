package user

import (
	"github.com/blackhorseya/lobster/internal/biz/user"
	"github.com/gin-gonic/gin"
)

type impl struct {
	biz user.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(biz user.IBiz) IHandler {
	return &impl{biz: biz}
}

// @Summary Signup
// @Description Signup
// @Tags Users
// @Accept application/json
// @Produce application/json
// @Param newUser body pb.Profile true "new user profile"
// @Success 201 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/users/signup [post]
func (i *impl) Signup(c *gin.Context) {
	// todo: 2021-03-01|17:04|doggy|implement me
	panic("implement me")
}

// @Summary Login
// @Description Login
// @Tags Users
// @Accept application/json
// @Produce application/json
// @Param user body pb.Profile true "user profile"
// @Success 201 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/users/login [post]
func (i *impl) Login(c *gin.Context) {
	// todo: 2021-03-01|17:04|doggy|implement me
	panic("implement me")
}
