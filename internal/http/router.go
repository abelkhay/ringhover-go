package http

import (
	"ringhover-go/internal/http/endpoints"
	"ringhover-go/internal/http/handlers"
	"ringhover-go/internal/logging"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *handlers.Handler) *gin.Engine {
	r := gin.New()
	r.Use(logging.ZapMiddleware(), gin.Recovery())

	api := r.Group(endpoints.APIBase)
	{
		api.GET(endpoints.HealthPath, func(c *gin.Context) { c.Status(200) })
		api.GET(endpoints.TasksPath, handler.GetRootTasksWithCategories) //1
		api.GET(endpoints.TaskSubtasks, handler.GetSubtasks)             //2
		api.POST(endpoints.TasksPath, handler.CreateTask)                //3
		api.PATCH(endpoints.TaskPath, handler.UpdateTask)                //4
		api.DELETE(endpoints.TaskPath, handler.DeleteTask)               //5
	}

	return r
}
