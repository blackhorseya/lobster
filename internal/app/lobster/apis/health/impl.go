package health

import (
	"net/http"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/health"
	"github.com/blackhorseya/lobster/internal/pkg/entities/errors"
	"github.com/blackhorseya/lobster/internal/pkg/entities/response"
	"github.com/blackhorseya/lobster/internal/pkg/utils/contextx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	biz    health.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(logger *zap.Logger, biz health.IBiz) IHandler {
	return &impl{
		logger: logger.With(zap.String("type", "HealthHandler")),
		biz:    biz,
	}
}

// Readiness @Summary Readiness
// @Description Show application was ready to start accepting traffic
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response
// @Failure 500 {object} errors.APPError
// @Router /readiness [get]
func (i *impl) Readiness(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	err := i.biz.Readiness(ctx)
	if err != nil {
		i.logger.Error(errors.ErrReadiness.Error(), zap.Error(err))
		c.Error(errors.ErrReadiness)
		return
	}

	c.JSON(http.StatusOK, response.OK)
}

// Liveness @Summary Liveness
// @Description to know when to restart an application
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response
// @Failure 500 {object} errors.APPError
// @Router /liveness [get]
func (i *impl) Liveness(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	err := i.biz.Liveness(ctx)
	if err != nil {
		i.logger.Error(errors.ErrReadiness.Error(), zap.Error(err))
		c.Error(errors.ErrLiveness)
		return
	}

	c.JSON(http.StatusOK, response.OK)
}
