package handlers

import (
	"net/http"
	"ringhover-go/internal/domain"
	"ringhover-go/internal/domain/req"
	"ringhover-go/internal/helpers"
	"ringhover-go/internal/http/httperr"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service domain.ModelisationServiceInterface
}

func NewTaskHandler(service domain.ModelisationServiceInterface) *Handler {
	return &Handler{service: service}
}

// GetSubtasks HTTP handler allows to get a list of substasks.
func (h *Handler) GetSubtasks(c *gin.Context) {
	taskId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || taskId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	rows, err := h.service.GetSubTasks(taskId)
	if err != nil {
		c.JSON(httperr.StatusCode(err), gin.H{"error": httperr.PublicMessage(err)})
		return
	}

	substasksForest := helpers.BuildSubtasksForest(rows, taskId)
	c.JSON(http.StatusOK, substasksForest)
}

// GetRootTasksWithCategory HTTP handler allows to get a list of all root tasks with their category.
func (h *Handler) GetRootTasksWithCategories(c *gin.Context) {

	tasks, err := h.service.GetRootTasks()
	if err != nil {
		c.JSON(httperr.StatusCode(err), gin.H{"error": httperr.PublicMessage(err)})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// CreateTask HTTP handler allows to create a task or a subtask.
func (h *Handler) CreateTask(c *gin.Context) {
	var request req.CreateTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
		return
	}

	request.Title = strings.TrimSpace(request.Title)
	if request.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}
	if len(request.Title) > 255 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title too long (max 255)"})
		return
	}

	task, err := h.service.CreateTask(request)
	if err != nil {
		c.JSON(httperr.StatusCode(err), gin.H{"error": httperr.PublicMessage(err)})
		return
	}
	c.JSON(http.StatusCreated, task)
}

// DeleteTask HTTP handler allows to delete a task and all associated substasks.
func (h *Handler) DeleteTask(c *gin.Context) {
	taskId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || taskId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err = h.service.DeleteTask(taskId); err != nil {
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// DeleteTask HTTP handler allows to delete a task and all associated substasks.
func (h *Handler) UpdateTask(c *gin.Context) {
	taskId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || taskId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var taskRequest req.UpdateTaskRequest
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
		return
	}

	taskUpdated, err := h.service.UpdateTask(taskId, taskRequest)
	if err != nil {
		c.JSON(httperr.StatusCode(err), gin.H{"error": httperr.PublicMessage(err)})
		return
	}
	c.JSON(http.StatusOK, taskUpdated)
}
