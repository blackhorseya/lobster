package goal

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/goal"
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/er"
	"github.com/blackhorseya/lobster/internal/pkg/entity/okr"
	"github.com/blackhorseya/lobster/internal/pkg/entity/response"
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

// GetByID @Summary Get a objective by id
// @Description Get a objective by id
// @Tags Goals
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of objective"
// @Success 200 {object} response.Response
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/goals/{id} [get]
func (i *impl) GetByID(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.Error(er.ErrInvalidID)
		return
	}

	ret, err := i.biz.GetByID(ctx, req.ID)
	if err != nil {
		i.logger.Error(er.ErrGetObjByID.Error(), zap.Error(err), zap.String("id", req.ID))
		c.Error(er.ErrGetObjByID)
		return
	}
	if ret == nil {
		i.logger.Error(er.ErrObjNotExists.Error(), zap.String("id", req.ID))
		c.Error(er.ErrObjNotExists)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

// List @Summary List all objectives
// @Description List all objectives
// @Tags Goals
// @Accept application/json
// @Produce application/json
// @Param page query integer false "page" default(1)
// @Param size query integer false "size of page" default(10)
// @Success 200 {object} response.Response
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/goals [get]
func (i *impl) List(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		i.logger.Error(er.ErrInvalidPage.Error(), zap.Error(err), zap.String("page", c.Query("page")))
		c.Error(er.ErrInvalidPage)
		return
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil {
		i.logger.Error(er.ErrInvalidSize.Error(), zap.Error(err), zap.String("size", c.Query("size")))
		c.Error(er.ErrInvalidSize)
		return
	}

	ret, err := i.biz.List(ctx, page, size)
	if err != nil {
		i.logger.Error(er.ErrListObj.Error(), zap.Error(err), zap.Int("page", page), zap.Int("size", size))
		c.Error(er.ErrListObj)
		return
	}
	if len(ret) == 0 {
		i.logger.Error(er.ErrObjNotExists.Error(), zap.Int("page", page), zap.Int("size", size))
		c.Error(er.ErrObjNotExists)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

// Create @Summary Create a objective
// @Description Create a objective
// @Tags Goals
// @Accept application/json
// @Produce application/json
// @Param created body okr.Goal true "created goal"
// @Success 201 {object} response.Response
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/goals [post]
func (i *impl) Create(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var created *okr.Goal
	if err := c.ShouldBindJSON(&created); err != nil {
		i.logger.Error(er.ErrCreateObj.Error(), zap.Error(err))
		c.Error(er.ErrCreateObj)
		return
	}

	if len(created.Title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error(), zap.Any("created", created))
		c.Error(er.ErrEmptyTitle)
		return
	}

	ret, err := i.biz.Create(ctx, created)
	if err != nil {
		i.logger.Error(er.ErrCreateObj.Error(), zap.Error(err), zap.Any("created", created))
		c.Error(er.ErrCreateObj)
		return
	}

	c.JSON(http.StatusCreated, response.OK.WithData(ret))
}

// ModifyTitle @Summary Modify title of goal
// @Description Modify title of goal
// @Tags Goals
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of goal"
// @Param updated body okr.Goal true "updated goal"
// @Success 200 {object} response.Response
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/goals/{id}/title [patch]
func (i *impl) ModifyTitle(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.Error(er.ErrInvalidID)
		return
	}

	var data *okr.Goal
	if err := c.ShouldBindJSON(&data); err != nil {
		i.logger.Error(er.ErrUpdateObj.Error(), zap.Error(err))
		c.Error(er.ErrUpdateObj)
		return
	}
	if len(data.Title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error(), zap.Any("data", data))
		c.Error(er.ErrEmptyTitle)
		return
	}

	ret, err := i.biz.ModifyTitle(ctx, req.ID, data.Title)
	if err != nil {
		i.logger.Error(er.ErrUpdateObj.Error(), zap.Error(err))
		c.Error(er.ErrUpdateObj)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

// Delete @Summary Get a objective by id
// @Description Get a objective by id
// @Tags Goals
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of objective"
// @Success 204 {object} string
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/goals/{id} [delete]
func (i *impl) Delete(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.Error(er.ErrInvalidID)
		return
	}

	err := i.biz.Delete(ctx, req.ID)
	if err != nil {
		i.logger.Error(er.ErrDeleteObj.Error(), zap.Error(err), zap.String("id", req.ID))
		c.Error(er.ErrDeleteObj)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
