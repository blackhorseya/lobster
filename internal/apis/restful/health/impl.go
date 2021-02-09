package health

import (
	"net/http"

	"github.com/blackhorseya/lobster/internal/biz/health"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	er "github.com/blackhorseya/lobster/internal/pkg/entities/error"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type impl struct {
	biz health.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(biz health.IBiz) IHandler {
	return &impl{biz: biz}
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
	ctx, ok := c.MustGet("ctx").(contextx.Contextx)
	if !ok {
		logrus.Error(er.ErrCTX)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.ErrCTX.Error(),
		})
		return
	}
	logger := ctx.WithField("func", "task getByID")

	err := i.biz.Readiness(ctx)
	if err != nil {
		logger.WithField("err", err).Error(er.ErrReadiness)
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
	ctx, ok := c.MustGet("ctx").(contextx.Contextx)
	if !ok {
		logrus.Error(er.ErrCTX)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.ErrCTX.Error(),
		})
		return
	}
	logger := ctx.WithField("func", "task getByID")

	err := i.biz.Liveness(ctx)
	if err != nil {
		logger.WithField("err", err).Error(er.ErrReadiness)
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.ErrReadiness.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
