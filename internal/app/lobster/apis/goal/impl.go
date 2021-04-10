package goal

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/goal"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	er "github.com/blackhorseya/lobster/internal/pkg/entities/error"
	"github.com/blackhorseya/lobster/internal/pkg/entities/okr"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	biz    goal.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(logger *zap.Logger, biz goal.IBiz) IHandler {
	return &impl{
		logger: logger.With(zap.String("type", "GoalHandler")),
		biz:    biz,
	}
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
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidID})
		return
	}

	ret, err := i.biz.GetByID(ctx, req.ID)
	if err != nil {
		i.logger.Error(er.ErrGetObjByID.Error(), zap.Error(err), zap.String("id", req.ID))
		c.JSON(http.StatusOK, gin.H{"error": er.ErrGetObjByID})
		return
	}
	if ret == nil {
		i.logger.Error(er.ErrObjectiveNotExists.Error(), zap.String("id", req.ID))
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
	ctx := c.MustGet("ctx").(contextx.Contextx)

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		i.logger.Error(er.ErrInvalidPage.Error(), zap.Error(err), zap.String("page", c.Query("page")))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidPage})
		return
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil {
		i.logger.Error(er.ErrInvalidSize.Error(), zap.Error(err), zap.String("size", c.Query("size")))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidSize})
		return
	}

	ret, err := i.biz.List(ctx, page, size)
	if err != nil {
		i.logger.Error(er.ErrListObjectives.Error(), zap.Error(err), zap.Int("page", page), zap.Int("size", size))
		c.JSON(http.StatusOK, gin.H{"error": er.ErrListObjectives})
		return
	}
	if len(ret) == 0 {
		i.logger.Error(er.ErrObjectiveNotExists.Error(), zap.Int("page", page), zap.Int("size", size))
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
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var created *okr.Goal
	if err := c.ShouldBindJSON(&created); err != nil {
		i.logger.Error(er.ErrCreateObjective.Error(), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.ErrCreateObjective})
		return
	}

	if len(created.Title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error(), zap.Any("created", created))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrEmptyTitle})
		return
	}

	ret, err := i.biz.Create(ctx, created)
	if err != nil {
		i.logger.Error(er.ErrCreateObjective.Error(), zap.Error(err), zap.Any("created", created))
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

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidID})
		return
	}

	var data *okr.Goal
	if err := c.ShouldBindJSON(&data); err != nil {
		i.logger.Error(er.ErrUpdateObj.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrUpdateObj})
		return
	}
	if len(data.Title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error(), zap.Any("data", data))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrEmptyTitle})
		return
	}

	ret, err := i.biz.ModifyTitle(ctx, req.ID, data.Title)
	if err != nil {
		i.logger.Error(er.ErrUpdateObj.Error(), zap.Error(err))
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
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidID})
		return
	}

	err := i.biz.Delete(ctx, req.ID)
	if err != nil {
		i.logger.Error(er.ErrDeleteObj.Error(), zap.Error(err), zap.String("id", req.ID))
		c.JSON(http.StatusOK, gin.H{"error": er.ErrDeleteObj})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
