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
		api.GET(endpoints.TasksPath, handler.GetRootTasksWithCategories)
		api.GET(endpoints.TaskSubtasks, handler.GetSubtasks)
		api.POST(endpoints.TasksPath, handler.CreateTask)
		api.PATCH(endpoints.TaskPath, handler.UpdateTask)
		api.DELETE(endpoints.TaskPath, handler.DeleteTask)
	}

	return r
}
