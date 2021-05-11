package task

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/task"
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/er"
	"github.com/blackhorseya/lobster/internal/pkg/entity/response"
	taskE "github.com/blackhorseya/lobster/internal/pkg/entity/task"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	biz    task.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(logger *zap.Logger, biz task.IBiz) IHandler {
	return &impl{
		logger: logger.With(zap.String("type", "TaskHandler")),
		biz:    biz,
	}
}

type reqID struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// GetByID @Summary Get a task by id
// @Description Get a task by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of task"
// @Success 200 {object} response.Response
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks/{id} [get]
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
		i.logger.Error(er.ErrGetTaskByID.Error(), zap.Error(err))
		c.Error(er.ErrGetTaskByID)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

// List @Summary List all tasks
// @Description List all tasks
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param page query integer false "page" default(1)
// @Param size query integer false "size of page" default(10)
// @Success 200 {object} response.Response
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks [get]
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
		i.logger.Error(er.ErrListTasks.Error(), zap.Error(err))
		c.Error(er.ErrListTasks)
		return
	}
	if len(ret) == 0 {
		i.logger.Error(er.ErrTaskNotExists.Error())
		c.Error(er.ErrTaskNotExists)
		return
	}

	c.JSON(http.StatusOK, ret)
}

// Create @Summary Create a task
// @Description Create a task
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param created body task.Task true "created task"
// @Success 201 {object} response.Response
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks [post]
func (i *impl) Create(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var data *taskE.Task
	if err := c.ShouldBindJSON(&data); err != nil {
		i.logger.Error(er.ErrCreateTask.Error(), zap.Error(err))
		c.Error(er.ErrCreateTask)
		return
	}

	if len(data.Title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error(), zap.Any("created", data))
		c.Error(er.ErrEmptyTitle)
		return
	}

	ret, err := i.biz.Create(ctx, data)
	if err != nil {
		i.logger.Error(er.ErrCreateTask.Error(), zap.Error(err), zap.Any("created", data))
		c.Error(er.ErrCreateTask)
		return
	}

	c.JSON(http.StatusCreated, response.OK.WithData(ret))
}

// UpdateStatus @Summary UpdateStatus a status of task by id
// @Description UpdateStatus a status of task by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of task"
// @Param updated body task.Task true "updated task"
// @Success 200 {object} response.Response
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks/{id}/status [patch]
func (i *impl) UpdateStatus(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.Error(er.ErrInvalidID)
		return
	}

	var data *taskE.Task
	if err := c.ShouldBindJSON(&data); err != nil {
		i.logger.Error(er.ErrCreateTask.Error(), zap.Error(err))
		c.Error(er.ErrUpdateTask)
		return
	}

	ret, err := i.biz.UpdateStatus(ctx, req.ID, data.Status)
	if err != nil {
		i.logger.Error(er.ErrUpdateTask.Error(), zap.Error(err))
		c.Error(er.ErrUpdateTask)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

// ModifyTitle @Summary ModifyTitle a title of task by id
// @Description ModifyTitle a status of task by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of task"
// @Param updated body task.Task true "updated task"
// @Success 200 {object} response.Response
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks/{id}/title [patch]
func (i *impl) ModifyTitle(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.Error(er.ErrInvalidID)
		return
	}

	var data *taskE.Task
	if err := c.ShouldBindJSON(&data); err != nil {
		i.logger.Error(er.ErrCreateTask.Error(), zap.Error(err))
		c.Error(er.ErrUpdateTask)
		return
	}
	if len(data.Title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error())
		c.Error(er.ErrEmptyTitle)
		return
	}

	ret, err := i.biz.ModifyTitle(ctx, req.ID, data.Title)
	if err != nil {
		i.logger.Error(er.ErrUpdateTask.Error(), zap.Error(err))
		c.Error(er.ErrUpdateTask)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

// Delete @Summary Delete a task by id
// @Description Delete a task by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of task"
// @Success 204 {object} string
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks/{id} [delete]
func (i *impl) Delete(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.Error(er.ErrInvalidID)
		return
	}

	if err := i.biz.Delete(ctx, req.ID); err != nil {
		i.logger.Error(er.ErrDeleteTask.Error(), zap.Error(err), zap.String("id", req.ID))
		c.Error(er.ErrDeleteTask)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
