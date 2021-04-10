package result

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/result"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	er "github.com/blackhorseya/lobster/internal/pkg/entities/error"
	"github.com/blackhorseya/lobster/internal/pkg/entities/okr"
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

// @Summary Get a key result by id
// @Description Get a key result by id
// @Tags Results
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of key result"
// @Success 200 {object} pb.Result
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /v1/results/{id} [get]
func (i *impl) GetByID(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ret, err := i.biz.GetByID(ctx, req.ID)
	if err != nil {
		i.logger.Error(er.ErrGetKRByID.Error(), zap.Error(err), zap.String("id", req.ID))
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.ErrGetKRByID})
		return
	}
	if ret == nil {
		i.logger.Error(er.ErrKRNotExists.Error(), zap.String("id", req.ID))
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, ret)
}

// @Summary Get key result by goal id
// @Description Get key result by goal id
// @Tags Results
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of goal"
// @Success 200 {array} pb.Result
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /v1/goals/{id}/results [get]
func (i *impl) GetByGoalID(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ret, err := i.biz.GetByGoalID(ctx, req.ID)
	if err != nil {
		i.logger.Error(er.ErrListKeyResult.Error(), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if len(ret) == 0 {
		i.logger.Error(er.ErrKRNotExists.Error())
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, ret)
}

// @Summary List all key results
// @Description List all key results
// @Tags Results
// @Accept application/json
// @Produce application/json
// @Param page query integer false "page" default(1)
// @Param size query integer false "size of page" default(10)
// @Success 200 {array} pb.Result
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /v1/results [get]
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
		i.logger.Error(er.ErrInvalidPage.Error(), zap.Error(err), zap.String("size", c.Query("size")))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidSize})
		return
	}

	ret, err := i.biz.List(ctx, page, size)
	if err != nil {
		i.logger.Error(er.ErrListKeyResult.Error(), zap.Error(err), zap.String("size", c.Query("size")))
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.ErrListKeyResult})
		return
	}
	if len(ret) == 0 {
		i.logger.Error(er.ErrKRNotExists.Error(), zap.String("size", c.Query("size")))
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, ret)
}

// @Summary Create a key result
// @Description Create a key result
// @Tags Results
// @Accept application/json
// @Produce application/json
// @Param created body pb.Result true "created key result"
// @Success 201 {object} pb.Result
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/results [post]
func (i *impl) Create(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var created *okr.Result
	err := c.ShouldBindJSON(&created)
	if err != nil {
		i.logger.Error(er.ErrCreateKR.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, er.ErrCreateKR)
		return
	}

	if len(created.Title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error(), zap.Any("created", created))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrEmptyTitle})
		return
	}
	_, err = uuid.Parse(created.GoalID)
	if err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err), zap.Any("created", created))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidID})
		return
	}

	ret, err := i.biz.LinkToGoal(ctx, created)
	if err != nil {
		i.logger.Error(er.ErrCreateKR.Error(), zap.Error(err), zap.Any("created", created))
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.ErrCreateKR})
		return
	}

	c.JSON(http.StatusCreated, ret)
}

// @Summary Modify title of result
// @Description Modify title of result
// @Tags Results
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of result"
// @Param updated body pb.Result true "updated result"
// @Success 200 {object} pb.Result
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/results/{id}/title [patch]
func (i *impl) ModifyTitle(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": er.ErrInvalidID})
		return
	}

	var data *okr.Result
	err := c.ShouldBindJSON(&data)
	if err != nil {
		i.logger.Error(er.ErrUpdateKeyResult.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, er.ErrUpdateKeyResult)
		return
	}
	if len(data.Title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error())
		c.JSON(http.StatusBadRequest, er.ErrEmptyTitle)
		return
	}

	ret, err := i.biz.ModifyTitle(ctx, req.ID, data.Title)
	if err != nil {
		i.logger.Error(er.ErrUpdateKeyResult.Error(), zap.Error(err))
		c.JSON(http.StatusInternalServerError, er.ErrUpdateKeyResult)
		return
	}

	c.JSON(http.StatusOK, ret)
}

// @Summary Delete a key result by id
// @Description Delete a key result by id
// @Tags Results
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of key result"
// @Success 204 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /v1/results/{id} [delete]
func (i *impl) Delete(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := i.biz.Delete(ctx, req.ID)
	if err != nil {
		i.logger.Error(er.ErrDeleteKeyResult.Error(), zap.Error(err), zap.String("id", req.ID))
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.ErrDeleteKeyResult})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
