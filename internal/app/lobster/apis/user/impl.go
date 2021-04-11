package user

import (
	"net/http"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/errors"
	"github.com/blackhorseya/lobster/internal/pkg/entities/response"
	user2 "github.com/blackhorseya/lobster/internal/pkg/entities/user"
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

// Signup @Summary Signup
// @Description Signup
// @Tags Users
// @Accept application/json
// @Produce application/json
// @Param newUser body user.Profile true "new user profile"
// @Success 201 {object} response.Response
// @Failure 400 {object} errors.APPError
// @Failure 500 {object} errors.APPError
// @Router /v1/users/signup [post]
func (i *impl) Signup(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var newUser *user2.Profile
	if err := c.ShouldBindJSON(&newUser); err != nil {
		i.logger.Error(errors.ErrSignup.Error())
		c.Error(errors.ErrSignup)
		return
	}

	ret, err := i.biz.Signup(ctx, newUser.Email, newUser.AccessToken)
	if err != nil {
		i.logger.Error(errors.ErrSignup.Error())
		c.Error(errors.ErrSignup)
		return
	}

	c.JSON(http.StatusCreated, response.OK.WithData(ret))
}

// Login @Summary Login
// @Description Login
// @Tags Users
// @Accept application/json
// @Produce application/json
// @Param user body user.Profile true "user profile"
// @Success 201 {object} response.Response
// @Failure 400 {object} errors.APPError
// @Failure 500 {object} errors.APPError
// @Router /v1/users/login [post]
func (i *impl) Login(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var data *user2.Profile
	if err := c.ShouldBindJSON(&data); err != nil {
		i.logger.Error(errors.ErrLogin.Error())
		c.Error(errors.ErrLogin)
		return
	}

	ret, err := i.biz.Login(ctx, data.Email, data.AccessToken)
	if err != nil {
		i.logger.Error(errors.ErrLogin.Error())
		c.Error(errors.ErrLogin)
		return
	}

	c.JSON(http.StatusCreated, response.OK.WithData(ret))
}
