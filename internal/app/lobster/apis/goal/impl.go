package goal

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/goal"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	er "github.com/blackhorseya/lobster/internal/pkg/entities/error"
	"github.com/blackhorseya/lobster/internal/pkg/pb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type impl struct {
	biz goal.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(biz goal.IBiz) IHandler {
	return &impl{biz: biz}
}

type reqID struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// @Summary Get a objective by id
// @Description Get a objective by id
// @Tags Goals
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of objective"
// @Success 200 {object} pb.Goal
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /v1/goals/{id} [get]
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
// @Tags Goals
// @Accept application/json
// @Produce application/json
// @Param page query integer false "page" default(1)
// @Param size query integer false "size of page" default(10)
// @Success 200 {array} pb.Goal
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /v1/goals [get]
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
// @Tags Goals
// @Accept application/json
// @Produce application/json
// @Param created body pb.Goal true "created goal"
// @Success 201 {object} pb.Goal
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/goals [post]
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

	var created *pb.Goal
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

// @Summary Modify title of goal
// @Description Modify title of goal
// @Tags Goals
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of goal"
// @Param updated body pb.Goal true "updated goal"
// @Success 200 {object} pb.Goal
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/goals/{id}/title [patch]
func (i *impl) ModifyTitle(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)
	logger := ctx.WithField("func", "ModifyTitle")

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		logger.WithField("err", err).Error(er.ErrInvalidID)
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidID})
		return
	}

	var data *pb.Goal
	if err := c.ShouldBindJSON(&data); err != nil {
		logger.WithField("err", err).Error(er.ErrUpdateObj)
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrUpdateObj})
		return
	}
	if len(data.Title) == 0 {
		logger.WithField("data", data).Error(er.ErrEmptyTitle)
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrEmptyTitle})
		return
	}

	ret, err := i.biz.ModifyTitle(ctx, req.ID, data.Title)
	if err != nil {
		logger.WithError(err).Error(er.ErrUpdateObj)
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.ErrUpdateObj})
		return
	}

	c.JSON(http.StatusOK, ret)
}

// @Summary Get a objective by id
// @Description Get a objective by id
// @Tags Goals
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of objective"
// @Success 204 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/goals/{id} [delete]
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
