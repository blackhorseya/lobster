package todo

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/lobster/internal/app/biz/todo"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	entities "github.com/blackhorseya/lobster/internal/pkg/entities/biz/todo"
	er "github.com/blackhorseya/lobster/internal/pkg/entities/error"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type impl struct {
	biz todo.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(biz todo.IBiz) IHandler {
	return &impl{biz: biz}
}

type reqID struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (i *impl) GetByID(c *gin.Context) {
	ctx, ok := c.MustGet("ctx").(contextx.Contextx)
	if !ok {
		logrus.Error(er.ErrCTX)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.ErrCTX.Error(),
		})
		return
	}
	logger := ctx.WithField("func", "task getByID")

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		logger.WithField("err", err).Error(er.ErrInvalidID)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if ret, err := i.biz.GetByID(ctx, req.ID); err != nil {
		logger.WithField("err", err).Error(er.ErrGetTaskByID)
		c.JSON(http.StatusOK, gin.H{"error": err})
		return
	} else {
		c.JSON(http.StatusOK, ret)
		return
	}
}

func (i *impl) List(c *gin.Context) {
	ctx, ok := c.MustGet("ctx").(contextx.Contextx)
	if !ok {
		logrus.Error(er.ErrCTX)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.ErrCTX.Error(),
		})
		return
	}
	logger := ctx.WithField("func", "task list")

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

	if ret, err := i.biz.List(ctx, page, size); err != nil {
		logger.WithField("error", err).Error(er.ErrListTasks)
		c.JSON(http.StatusOK, gin.H{"error": er.ErrListTasks})
		return
	} else {
		c.JSON(http.StatusOK, ret)
		return
	}
}

func (i *impl) Create(c *gin.Context) {
	ctx, ok := c.MustGet("ctx").(contextx.Contextx)
	if !ok {
		logrus.Error(er.ErrCTX)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.ErrCTX.Error(),
		})
		return
	}
	logger := ctx.WithField("func", "task list")

	var task *entities.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		logger.WithField("error", err).Error(er.ErrCreateTask)
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrCreateTask})
		return
	}

	if len(task.Title) == 0 {
		logger.WithField("task", task).Error(er.ErrEmptyTitle)
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrEmptyTitle})
		return
	}

	if ret, err := i.biz.Create(ctx, task); err != nil {
		logger.WithFields(logrus.Fields{"error": err, "task": task}).Error(er.ErrCreateTask)
		c.JSON(http.StatusOK, gin.H{"error": er.ErrCreateTask})
		return
	} else {
		c.JSON(http.StatusOK, ret)
		return
	}
}

func (i *impl) Update(c *gin.Context) {
	ctx, ok := c.MustGet("ctx").(contextx.Contextx)
	if !ok {
		logrus.Error(er.ErrCTX)
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.ErrCTX.Error()})
		return
	}
	logger := ctx.WithField("func", "task list")

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		logger.WithField("err", err).Error(er.ErrInvalidID)
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidID})
		return
	}

	var task *entities.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		logger.WithField("error", err).Error(er.ErrCreateTask)
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrCreateTask})
		return
	}

	if len(task.Title) == 0 {
		logger.WithField("task", task).Error(er.ErrEmptyTitle)
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrEmptyTitle})
		return
	}

	// todo: 2021-01-15|10:37|doggy|implement me
	panic("implement me")
}

func (i *impl) Delete(c *gin.Context) {
	ctx, ok := c.MustGet("ctx").(contextx.Contextx)
	if !ok {
		logrus.Error(er.ErrCTX)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.ErrCTX.Error(),
		})
		return
	}
	logger := ctx.WithField("func", "task list")

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		logger.WithField("err", err).Error(er.ErrInvalidID)
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidID})
		return
	}

	if err := i.biz.Delete(ctx, req.ID); err != nil {
		logger.WithFields(logrus.Fields{"error": err, "id": req.ID}).Error(er.ErrDeleteTask)
		c.JSON(http.StatusOK, gin.H{"error": er.ErrDeleteTask})
		return
	} else {
		c.JSON(http.StatusNoContent, nil)
		return
	}
}
