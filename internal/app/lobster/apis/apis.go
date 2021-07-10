package apis

import (
	// import swagger docs
	_ "github.com/blackhorseya/lobster/api/docs"
	"github.com/blackhorseya/lobster/internal/app/lobster/apis/health"
	"github.com/blackhorseya/lobster/internal/app/lobster/apis/task"
	"github.com/blackhorseya/lobster/internal/pkg/infra/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CreateInitHandlerFn serve caller to create init handler
func CreateInitHandlerFn(
	health health.IHandler,
	taskH task.IHandler) http.InitHandlers {
	return func(r *gin.Engine) {
		api := r.Group("api")
		{
			api.GET("readiness", health.Readiness)
			api.GET("liveness", health.Liveness)

			// open any environments can access swagger
			api.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

			v1 := api.Group("v1")
			{
				tasks := v1.Group("tasks")
				{
					tasks.GET("", taskH.List)
					tasks.GET(":id", taskH.GetByID)
					tasks.POST("", taskH.Create)
					tasks.DELETE(":id", taskH.Delete)
					tasks.PATCH(":id/status", taskH.UpdateStatus)
					tasks.PATCH(":id/title", taskH.ModifyTitle)
				}
			}
		}
	}
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	task.ProviderSet,
	CreateInitHandlerFn,
)
