package health

import (
	"net/http"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/health"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	er "github.com/blackhorseya/lobster/internal/pkg/entities/error"
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

// @Summary Readiness
// @Description Show application was ready to start accepting traffic
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /readiness [get]
func (i *impl) Readiness(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	err := i.biz.Readiness(ctx)
	if err != nil {
		i.logger.Error(er.ErrReadiness.Error(), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.ErrReadiness.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

// @Summary Liveness
// @Description to know when to restart an application
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /liveness [get]
func (i *impl) Liveness(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	err := i.biz.Liveness(ctx)
	if err != nil {
		i.logger.Error(er.ErrReadiness.Error(), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.ErrReadiness.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
