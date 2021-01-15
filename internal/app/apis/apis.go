package apis

import (
	// import swagger docs
	_ "github.com/blackhorseya/lobster/internal/app/apis/docs"
	"github.com/blackhorseya/lobster/internal/app/apis/health"
	"github.com/blackhorseya/lobster/internal/app/apis/todo"
	"github.com/blackhorseya/lobster/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CreateInitHandlerFn serve caller to create init handler
func CreateInitHandlerFn(health health.IHandler, todoHandler todo.IHandler) http.InitHandlers {
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
			}
		}
	}
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	todo.ProviderSet,
	CreateInitHandlerFn,
)
