package result

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/result"
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/er"
	"github.com/blackhorseya/lobster/internal/pkg/entity/okr"
	"github.com/blackhorseya/lobster/internal/pkg/entity/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	biz    result.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(logger *zap.Logger, biz result.IBiz) IHandler {
	return &impl{
		logger: logger.With(zap.String("type", "ResultHandler")),
		biz:    biz,
	}
}

type reqID struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// GetByID @Summary Get a key result by id
// @Description Get a key result by id
// @Tags Results
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of key result"
// @Success 200 {object} response.Response
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/results/{id} [get]
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
		i.logger.Error(er.ErrGetKRByID.Error(), zap.Error(err), zap.String("id", req.ID))
		c.Error(er.ErrGetKRByID)
		return
	}
	if ret == nil {
		i.logger.Error(er.ErrKRNotExists.Error(), zap.String("id", req.ID))
		c.Error(er.ErrKRNotExists)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

// GetByGoalID @Summary Get key result by goal id
// @Description Get key result by goal id
// @Tags Results
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of goal"
// @Success 200 {object} response.Response
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/goals/{id}/results [get]
func (i *impl) GetByGoalID(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.Error(er.ErrInvalidID)
		return
	}

	ret, err := i.biz.GetByGoalID(ctx, req.ID)
	if err != nil {
		i.logger.Error(er.ErrListKR.Error(), zap.Error(err))
		c.Error(er.ErrListKR)
		return
	}
	if len(ret) == 0 {
		i.logger.Error(er.ErrKRNotExists.Error())
		c.Error(er.ErrKRNotExists)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

// List @Summary List all key results
// @Description List all key results
// @Tags Results
// @Accept application/json
// @Produce application/json
// @Param page query integer false "page" default(1)
// @Param size query integer false "size of page" default(10)
// @Success 200 {object} response.Response
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/results [get]
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
		i.logger.Error(er.ErrInvalidPage.Error(), zap.Error(err), zap.String("size", c.Query("size")))
		c.Error(er.ErrInvalidSize)
		return
	}

	ret, err := i.biz.List(ctx, page, size)
	if err != nil {
		i.logger.Error(er.ErrListKR.Error(), zap.Error(err), zap.String("size", c.Query("size")))
		c.Error(er.ErrListKR)
		return
	}
	if len(ret) == 0 {
		i.logger.Error(er.ErrKRNotExists.Error(), zap.String("size", c.Query("size")))
		c.Error(er.ErrKRNotExists)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

// Create @Summary Create a key result
// @Description Create a key result
// @Tags Results
// @Accept application/json
// @Produce application/json
// @Param created body okr.Result true "created key result"
// @Success 201 {object} response.Response
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/results [post]
func (i *impl) Create(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var created *okr.Result
	err := c.ShouldBindJSON(&created)
	if err != nil {
		i.logger.Error(er.ErrCreateKR.Error(), zap.Error(err))
		c.Error(er.ErrCreateKR)
		return
	}

	if len(created.Title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error(), zap.Any("created", created))
		c.Error(er.ErrEmptyTitle)
		return
	}
	_, err = uuid.Parse(created.GoalID)
	if err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err), zap.Any("created", created))
		c.Error(er.ErrInvalidID)
		return
	}

	ret, err := i.biz.LinkToGoal(ctx, created)
	if err != nil {
		i.logger.Error(er.ErrCreateKR.Error(), zap.Error(err), zap.Any("created", created))
		c.Error(er.ErrCreateKR)
		return
	}

	c.JSON(http.StatusCreated, response.OK.WithData(ret))
}

// ModifyTitle @Summary Modify title of result
// @Description Modify title of result
// @Tags Results
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of result"
// @Param updated body okr.Result true "updated result"
// @Success 200 {object} response.Response
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/results/{id}/title [patch]
func (i *impl) ModifyTitle(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.Error(er.ErrInvalidID)
		return
	}

	var data *okr.Result
	err := c.ShouldBindJSON(&data)
	if err != nil {
		i.logger.Error(er.ErrUpdateKR.Error(), zap.Error(err))
		c.Error(er.ErrUpdateKR)
		return
	}
	if len(data.Title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error())
		c.Error(er.ErrEmptyTitle)
		return
	}

	ret, err := i.biz.ModifyTitle(ctx, req.ID, data.Title)
	if err != nil {
		i.logger.Error(er.ErrUpdateKR.Error(), zap.Error(err))
		c.Error(er.ErrUpdateKR)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

// Delete @Summary Delete a key result by id
// @Description Delete a key result by id
// @Tags Results
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of key result"
// @Success 204 {object} string
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/results/{id} [delete]
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
		i.logger.Error(er.ErrDeleteKR.Error(), zap.Error(err), zap.String("id", req.ID))
		c.Error(er.ErrDeleteKR)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
