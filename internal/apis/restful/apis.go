package restful

import (
	// import swagger docs
	_ "github.com/blackhorseya/lobster/internal/apis/restful/docs"
	"github.com/blackhorseya/lobster/internal/apis/restful/goal"
	"github.com/blackhorseya/lobster/internal/apis/restful/health"
	"github.com/blackhorseya/lobster/internal/apis/restful/kr"
	"github.com/blackhorseya/lobster/internal/apis/restful/task"
	"github.com/blackhorseya/lobster/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CreateInitHandlerFn serve caller to create init handler
func CreateInitHandlerFn(health health.IHandler, taskH task.IHandler, goalH goal.IHandler, krH kr.IHandler) http.InitHandlers {
	return func(r *gin.Engine) {
		api := r.Group("api")
		{
			api.GET("readiness", health.Readiness)
			api.GET("liveness", health.Liveness)

			if mode := gin.Mode(); mode != gin.ReleaseMode {
				api.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
			}

			v1 := api.Group("v1")
			{
				tasks := v1.Group("tasks")
				{
					tasks.GET("", taskH.List)
					tasks.GET(":id", taskH.GetByID)
					tasks.POST("", taskH.Create)
					tasks.PUT(":id", taskH.Update)
					tasks.DELETE(":id", taskH.Delete)
				}

				goals := v1.Group("goals")
				{
					goals.GET("", goalH.List)
					goals.GET(":id", goalH.GetByID)
					goals.POST("", goalH.Create)
					goals.PUT(":id", goalH.Update)
					goals.DELETE(":id", goalH.Delete)
				}

				krs := v1.Group("krs")
				{
					krs.GET("", krH.List)
					krs.GET(":id", krH.GetByID)
					krs.POST("", krH.Create)
					krs.PUT(":id", krH.Update)
					krs.DELETE(":id", krH.Delete)
				}
			}
		}
	}
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	task.ProviderSet,
	goal.ProviderSet,
	kr.ProviderSet,
	CreateInitHandlerFn,
)
