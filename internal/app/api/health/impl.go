package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type impl struct {
}

// NewImpl serve caller to create an IHandler
func NewImpl() IHandler {
	return &impl{}
}

func (i *impl) Readiness(c *gin.Context) {
	// todo: 2021-01-12|10:12|doggy|implement me
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func (i *impl) Liveness(c *gin.Context) {
	// todo: 2021-01-12|10:12|doggy|implement me
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
