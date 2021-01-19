package objective

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/lobster/internal/app/biz/objective"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	er "github.com/blackhorseya/lobster/internal/pkg/entities/error"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type impl struct {
	biz objective.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(biz objective.IBiz) IHandler {
	return &impl{biz: biz}
}

type reqID struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// @Summary Get a objective by id
// @Description Get a objective by id
// @Tags Objectives
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of objective"
// @Success 200 {object} okr.Objective
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /v1/objectives/{id} [get]
func (i *impl) GetByID(c *gin.Context) {
	ctx, ok := c.MustGet("ctx").(contextx.Contextx)
	if !ok {
		logrus.Error(er.ErrCTX)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.ErrCTX.Error(),
		})
		return
	}
	logger := ctx.WithField("func", "objective getByID")

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		logger.WithField("err", err).Error(er.ErrInvalidID)
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidID})
		return
	}

	ret, err := i.biz.GetByID(ctx, req.ID)
	if err != nil {
		logger.WithFields(logrus.Fields{"err": err, "id": req.ID}).Error(er.ErrGetObjByID)
		c.JSON(http.StatusOK, gin.H{"error": er.ErrGetObjByID})
		return
	}
	if ret == nil {
		logger.WithFields(logrus.Fields{"id": req.ID}).Error(er.ErrObjectiveNotExists)
		c.JSON(http.StatusNotFound, gin.H{"error": er.ErrObjectiveNotExists})
		return
	}

	c.JSON(http.StatusOK, ret)
}

// @Summary List all objectives
// @Description List all objectives
// @Tags Objectives
// @Accept application/json
// @Produce application/json
// @Param page query integer false "page" default(1)
// @Param size query integer false "size of page" default(10)
// @Success 200 {array} okr.Objective
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /v1/objectives [get]
func (i *impl) List(c *gin.Context) {
	ctx, ok := c.MustGet("ctx").(contextx.Contextx)
	if !ok {
		logrus.Error(er.ErrCTX)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.ErrCTX.Error(),
		})
		return
	}
	logger := ctx.WithField("func", "objective list")

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

	ret, err := i.biz.List(ctx, page, size)
	if err != nil {
		logger.WithFields(logrus.Fields{"err": err, "page": page, "size": size}).Error(er.ErrListObjectives)
		c.JSON(http.StatusOK, gin.H{"error": er.ErrListObjectives})
		return
	}
	if len(ret) == 0 {
		logger.WithFields(logrus.Fields{"page": page, "size": size}).Error(er.ErrObjectiveNotExists)
		c.JSON(http.StatusNotFound, gin.H{"error": er.ErrObjectiveNotExists})
		return
	}

	c.JSON(http.StatusOK, ret)
}

// @Summary Create a objective
// @Description Create a objective
// @Tags Objectives
// @Accept application/json
// @Produce application/json
// @Param created body okr.Objective true "created objective"
// @Success 201 {object} okr.Objective
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/objectives [post]
func (i *impl) Create(c *gin.Context) {
	ctx, ok := c.MustGet("ctx").(contextx.Contextx)
	if !ok {
		logrus.Error(er.ErrCTX)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.ErrCTX.Error(),
		})
		return
	}
	logger := ctx.WithField("func", "objective create")

	var created *okr.Objective
	if err := c.ShouldBindJSON(&created); err != nil {
		logger.WithField("err", err).Error(er.ErrCreateObjective)
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.ErrCreateObjective})
		return
	}

	if len(created.Title) == 0 {
		logger.WithField("create", created).Error(er.ErrEmptyTitle)
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrEmptyTitle})
		return
	}

	ret, err := i.biz.Create(ctx, created)
	if err != nil {
		logger.WithFields(logrus.Fields{"err": err, "created": created}).Error(er.ErrCreateObjective)
		c.JSON(http.StatusOK, gin.H{"error": er.ErrCreateObjective})
		return
	}

	c.JSON(http.StatusCreated, ret)
}

// @Summary Update a objective by id
// @Description Update a objective by id
// @Tags Objectives
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of objective"
// @Param created body okr.Objective true "created objective"
// @Success 200 {object} okr.Objective
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/objectives/{id} [put]
func (i *impl) Update(c *gin.Context) {
	ctx, ok := c.MustGet("ctx").(contextx.Contextx)
	if !ok {
		logrus.Error(er.ErrCTX)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.ErrCTX.Error(),
		})
		return
	}
	logger := ctx.WithField("func", "objective update")

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		logger.WithField("err", err).Error(er.ErrInvalidID)
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidID})
		return
	}

	var updated *okr.Objective
	if err := c.ShouldBindJSON(&updated); err != nil {
		logger.WithField("err", err).Error(er.ErrCreateObjective)
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.ErrCreateObjective})
		return
	}

	if len(updated.Title) == 0 {
		logger.WithField("updated", updated).Error(er.ErrEmptyTitle)
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrEmptyTitle})
		return
	}

	ret, err := i.biz.Update(ctx, updated)
	if err != nil {
		logger.WithFields(logrus.Fields{"err": err, "updated": updated}).Error(er.ErrUpdateObj)
		c.JSON(http.StatusOK, gin.H{"error": er.ErrUpdateObj})
		return
	}

	c.JSON(http.StatusOK, ret)
}

// @Summary Get a objective by id
// @Description Get a objective by id
// @Tags Objectives
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of objective"
// @Success 204 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/objectives/{id} [delete]
func (i *impl) Delete(c *gin.Context) {
	ctx, ok := c.MustGet("ctx").(contextx.Contextx)
	if !ok {
		logrus.Error(er.ErrCTX)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.ErrCTX.Error(),
		})
		return
	}
	logger := ctx.WithField("func", "objective update")

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		logger.WithField("err", err).Error(er.ErrInvalidID)
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidID})
		return
	}

	err := i.biz.Delete(ctx, req.ID)
	if err != nil {
		logger.WithFields(logrus.Fields{"err": err, "id": req.ID}).Error(er.ErrDeleteObj)
		c.JSON(http.StatusOK, gin.H{"error": er.ErrDeleteObj})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
