package restful

import (
	// import swagger docs
	_ "github.com/blackhorseya/lobster/internal/apis/restful/docs"
	"github.com/blackhorseya/lobster/internal/apis/restful/goal"
	"github.com/blackhorseya/lobster/internal/apis/restful/health"
	"github.com/blackhorseya/lobster/internal/apis/restful/result"
	"github.com/blackhorseya/lobster/internal/apis/restful/task"
	"github.com/blackhorseya/lobster/internal/apis/restful/user"
	"github.com/blackhorseya/lobster/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CreateInitHandlerFn serve caller to create init handler
func CreateInitHandlerFn(
	health health.IHandler,
	taskH task.IHandler,
	goalH goal.IHandler,
	resultH result.IHandler,
	userH user.IHandler) http.InitHandlers {
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

				goals := v1.Group("goals")
				{
					goals.GET("", goalH.List)
					goals.GET(":id", goalH.GetByID)
					goals.POST("", goalH.Create)
					goals.DELETE(":id", goalH.Delete)
					goals.GET(":id/results", resultH.GetByGoalID)
				}

				results := v1.Group("results")
				{
					results.GET("", resultH.List)
					results.GET(":id", resultH.GetByID)
					results.POST("", resultH.Create)
					results.PATCH(":id/title", resultH.ModifyTitle)
					results.DELETE(":id", resultH.Delete)
				}

				users := v1.Group("users")
				{
					users.POST("signup", userH.Signup)
					users.POST("login", userH.Login)
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
	result.ProviderSet,
	user.ProviderSet,
	CreateInitHandlerFn,
)
