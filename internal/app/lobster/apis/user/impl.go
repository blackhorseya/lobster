package user

import (
	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	biz    user.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(logger *zap.Logger, biz user.IBiz) IHandler {
	return &impl{
		logger: logger.With(zap.String("type", "UserHandler")),
		biz:    biz,
	}
}

// GetByID
// @Summary Get an user by id
// @Description Get an user by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path integer true "ID of user"
// @Success 200 {object} response.Response{data=user.Profile}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/users/{id} [get]
func (i *impl) GetByID(c *gin.Context) {
	// todo: 2021-07-11|08:53|Sean|implement me
	panic("implement me")
}

// Signup
// @Summary Signup
// @Description Signup
// @Tags Auth
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param email formData string true "email"
// @Param password formData string true "password"
// @Success 201 {object} response.Response{data=user.Profile}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/auth/signup [post]
func (i *impl) Signup(c *gin.Context) {
	// todo: 2021-07-11|08:53|Sean|implement me
	panic("implement me")
}

// Login
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param email formData string true "email"
// @Param password formData string true "password"
// @Success 201 {object} response.Response{data=user.Profile}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/auth/login [post]
func (i *impl) Login(c *gin.Context) {
	// todo: 2021-07-11|08:53|Sean|implement me
	panic("implement me")
}
