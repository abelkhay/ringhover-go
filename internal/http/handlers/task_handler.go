package handlers

import (
	"net/http"
	"ringhover-go/internal/domain"
	"ringhover-go/internal/helpers"
	"ringhover-go/internal/http/httperr"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service domain.ModelisationServiceInterface
}

func NewTaskHandler(service domain.ModelisationServiceInterface) *Handler {
	return &Handler{service: service}
}

// GetSubtasks HTTP handler allows to get a list of substasks
func (h *Handler) GetSubtasks(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	rows, err := h.service.GetSubTasks(id)
	if err != nil {
		c.JSON(httperr.StatusCode(err), gin.H{"error": httperr.PublicMessage(err)})
		return
	}

	substasksForest := helpers.BuildSubtasksForest(rows, id)
	c.JSON(http.StatusOK, substasksForest)
}
