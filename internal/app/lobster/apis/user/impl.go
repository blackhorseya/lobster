package user

import (
	"fmt"
	"net/http"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	user2 "github.com/blackhorseya/lobster/internal/pkg/entities/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	// ErrSignup means user signup is failure
	ErrSignup = fmt.Errorf("user signup is failure")

	// ErrLogin means user login is failure
	ErrLogin = fmt.Errorf("user login is failure")
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
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var newUser *user2.Profile
	if err := c.ShouldBindJSON(&newUser); err != nil {
		i.logger.Error(ErrSignup.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrSignup})
		return
	}

	ret, err := i.biz.Signup(ctx, newUser.Email, newUser.AccessToken)
	if err != nil {
		i.logger.Error(ErrSignup.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrSignup})
		return
	}

	c.JSON(http.StatusCreated, ret)
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
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var data *user2.Profile
	if err := c.ShouldBindJSON(&data); err != nil {
		i.logger.Error(ErrLogin.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrLogin})
		return
	}

	ret, err := i.biz.Login(ctx, data.Email, data.AccessToken)
	if err != nil {
		i.logger.Error(ErrLogin.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrLogin})
		return
	}

	c.JSON(http.StatusCreated, ret)
}
