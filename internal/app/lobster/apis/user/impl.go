package user

import (
	"net/http"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user"
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/er"
	"github.com/blackhorseya/lobster/internal/pkg/entity/response"
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

type reqID struct {
	ID int64 `uri:"id" binding:"required"`
}

// GetByID
// @Summary Get an user by id
// @Description Get an user by id
// @Tags UsersD
// @Accept application/json
// @Produce application/json
// @Param id path integer true "ID of user"
// @Success 200 {object} response.Response{data=user.Profile}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/users/{id} [get]
func (i *impl) GetByID(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.Error(er.ErrInvalidID)
		return
	}

	ret, err := i.biz.GetByID(ctx, req.ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
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
	ctx := c.MustGet("ctx").(contextx.Contextx)

	email := c.PostForm("email")
	password := c.PostForm("password")

	ret, err := i.biz.Signup(ctx, email, password)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response.OK.WithData(ret))
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
	ctx := c.MustGet("ctx").(contextx.Contextx)

	email := c.PostForm("email")
	password := c.PostForm("password")

	ret, err := i.biz.Login(ctx, email, password)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response.OK.WithData(ret))
}
