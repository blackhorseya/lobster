package task

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/task"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	er "github.com/blackhorseya/lobster/internal/pkg/entities/error"
	taskE "github.com/blackhorseya/lobster/internal/pkg/entities/task"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

// @Summary Get a task by id
// @Description Get a task by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of task"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/tasks/{id} [get]
func (i *impl) GetByID(c *gin.Context) {
	ctx, ok := c.MustGet("ctx").(contextx.Contextx)
	if !ok {
		logrus.Error(er.ErrCTX)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.ErrCTX.Error(),
		})
		return
	}

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ret, err := i.biz.GetByID(ctx, req.ID)
	if err != nil {
		i.logger.Error(er.ErrGetTaskByID.Error(), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, ret)
	return
}

// @Summary List all tasks
// @Description List all tasks
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param page query integer false "page" default(1)
// @Param size query integer false "size of page" default(10)
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/tasks [get]
func (i *impl) List(c *gin.Context) {
	ctx, ok := c.MustGet("ctx").(contextx.Contextx)
	if !ok {
		logrus.Error(er.ErrCTX)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.ErrCTX.Error(),
		})
		return
	}

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
		i.logger.Error(er.ErrListTasks.Error(), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{"error": er.ErrListTasks})
		return
	}
	if len(ret) == 0 {
		i.logger.Error(er.ErrTaskNotExists.Error())
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, ret)
	return
}

// @Summary Create a task
// @Description Create a task
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param created body pb.Task true "created task"
// @Success 200 {object} string
// @Success 201 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/tasks [post]
func (i *impl) Create(c *gin.Context) {
	ctx, ok := c.MustGet("ctx").(contextx.Contextx)
	if !ok {
		logrus.Error(er.ErrCTX)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.ErrCTX.Error(),
		})
		return
	}

	var task *taskE.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		i.logger.Error(er.ErrCreateTask.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrCreateTask})
		return
	}

	if len(task.Title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error(), zap.Any("task", task))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrEmptyTitle})
		return
	}

	ret, err := i.biz.Create(ctx, task)
	if err != nil {
		i.logger.Error(er.ErrCreateTask.Error(), zap.Error(err), zap.Any("task", task))
		c.JSON(http.StatusOK, gin.H{"error": er.ErrCreateTask})
		return
	}

	c.JSON(http.StatusCreated, ret)
	return
}

// @Summary UpdateStatus a status of task by id
// @Description UpdateStatus a status of task by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of task"
// @Param updated body pb.Task true "updated task"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/tasks/{id}/status [patch]
func (i *impl) UpdateStatus(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidID})
		return
	}

	var data *taskE.Task
	if err := c.ShouldBindJSON(&data); err != nil {
		i.logger.Error(er.ErrCreateTask.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrCreateTask})
		return
	}

	ret, err := i.biz.UpdateStatus(ctx, req.ID, data.Status)
	if err != nil {
		i.logger.Error(er.ErrUpdateTask.Error(), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.ErrUpdateTask})
		return
	}

	c.JSON(http.StatusOK, ret)
}

// @Summary ModifyTitle a title of task by id
// @Description ModifyTitle a status of task by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of task"
// @Param updated body pb.Task true "updated task"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/tasks/{id}/title [patch]
func (i *impl) ModifyTitle(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidID})
		return
	}

	var data *taskE.Task
	if err := c.ShouldBindJSON(&data); err != nil {
		i.logger.Error(er.ErrCreateTask.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrCreateTask})
		return
	}
	if len(data.Title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrEmptyTitle})
		return
	}

	ret, err := i.biz.ModifyTitle(ctx, req.ID, data.Title)
	if err != nil {
		i.logger.Error(er.ErrUpdateTask.Error(), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.ErrUpdateTask})
		return
	}

	c.JSON(http.StatusOK, ret)
}

// @Summary Delete a task by id
// @Description Delete a task by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of task"
// @Success 200 {object} string
// @Success 204 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/tasks/{id} [delete]
func (i *impl) Delete(c *gin.Context) {
	ctx, ok := c.MustGet("ctx").(contextx.Contextx)
	if !ok {
		logrus.Error(er.ErrCTX)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.ErrCTX.Error(),
		})
		return
	}

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidID})
		return
	}

	if err := i.biz.Delete(ctx, req.ID); err != nil {
		i.logger.Error(er.ErrDeleteTask.Error(), zap.Error(err), zap.String("id", req.ID))
		c.JSON(http.StatusOK, gin.H{"error": er.ErrDeleteTask})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
