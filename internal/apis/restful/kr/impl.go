package kr

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/lobster/internal/biz/kr"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	er "github.com/blackhorseya/lobster/internal/pkg/entities/error"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

// @Summary List all key results
// @Description List all key results
// @Tags KeyResults
// @Accept application/json
// @Produce application/json
// @Param page query integer false "page" default(1)
// @Param size query integer false "size of page" default(10)
// @Success 200 {array} okr.KeyResult
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /v1/krs [get]
func (i *impl) List(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)
	logger := ctx.WithField("func", "List")

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		logger.WithFields(logrus.Fields{"error": err, "page": c.Query("page")}).Error(er.ErrInvalidPage)
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidPage})
		return
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil {
		logger.WithFields(logrus.Fields{"error": err, "size": c.Query("size")}).Error(er.ErrInvalidSize)
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidSize})
		return
	}

	logger = logger.WithFields(logrus.Fields{"page": page, "size": size})

	ret, err := i.biz.List(ctx, page, size)
	if err != nil {
		logger.WithError(err).Error(er.ErrListKeyResult)
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.ErrListKeyResult})
		return
	}
	if len(ret) == 0 {
		logger.Error(er.ErrKRNotExists)
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, ret)
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
