package kr

import (
	"net/http"

	"github.com/blackhorseya/lobster/internal/biz/kr"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	er "github.com/blackhorseya/lobster/internal/pkg/entities/error"
	"github.com/gin-gonic/gin"
)

type impl struct {
	biz kr.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(biz kr.IBiz) IHandler {
	return &impl{biz: biz}
}

type reqID struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// @Summary Get a key result by id
// @Description Get a key result by id
// @Tags KeyResults
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of key result"
// @Success 200 {object} okr.KeyResult
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /v1/krs/{id} [get]
func (i *impl) GetByID(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)
	logger := ctx.WithField("func", "GetByID")

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		logger.WithField("err", err).Error(er.ErrInvalidID)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ret, err := i.biz.GetByID(ctx, req.ID)
	if err != nil {
		logger.WithError(err).WithField("id", req.ID).Error(er.ErrGetKRByID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.ErrGetKRByID})
		return
	}
	if ret == nil {
		logger.WithField("id", req.ID).Error(er.ErrKRNotExists)
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, ret)
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
