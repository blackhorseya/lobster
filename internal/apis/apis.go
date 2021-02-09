package apis

import (
	// import swagger docs
	_ "github.com/blackhorseya/lobster/internal/apis/restful/docs"
	"github.com/blackhorseya/lobster/internal/apis/restful/health"
	"github.com/blackhorseya/lobster/internal/apis/restful/objective"
	"github.com/blackhorseya/lobster/internal/apis/restful/todo"
	"github.com/blackhorseya/lobster/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CreateInitHandlerFn serve caller to create init handler
func CreateInitHandlerFn(health health.IHandler, todoHandler todo.IHandler, objHandler objective.IHandler) http.InitHandlers {
	return func(r *gin.Engine) {
		api := r.Group("/api")
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
					tasks.GET("", todoHandler.List)
					tasks.GET(":id", todoHandler.GetByID)
					tasks.POST("", todoHandler.Create)
					tasks.PUT(":id", todoHandler.Update)
					tasks.DELETE(":id", todoHandler.Delete)
				}

				objs := v1.Group("objectives")
				{
					objs.GET("", objHandler.List)
					objs.GET(":id", objHandler.GetByID)
					objs.POST("", objHandler.Create)
					objs.PUT(":id", objHandler.Update)
					objs.DELETE(":id", objHandler.Delete)
				}
			}
		}
	}
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	todo.ProviderSet,
	objective.ProviderSet,
	CreateInitHandlerFn,
)
