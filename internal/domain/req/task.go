package req

import (
	"ringhover-go/internal/domain/models"
	"time"
)

// CreateTask request struct onto which client requests should be successfully mapped.
type CreateTask struct {
	Title        string         `json:"title" binding:"required"`
	Description  *string        `json:"description"`
	Status       *models.Status `json:"status"`
	Priority     *int           `json:"priority"`
	DueDate      *time.Time     `json:"due_date"`
	ParentTaskID *uint64        `json:"parent_task_id"`
	CategoryID   *uint64        `json:"category_id"`
}

// UpdateTask request struct onto which client requests should be successfully mapped.
type UpdateTask struct {
	Title       *string        `json:"title"`
	Description *string        `json:"description"`
	Status      *models.Status `json:"status"`
	Priority    *int           `json:"priority"`
	DueDate     *time.Time     `json:"due_date"`
	CompletedAt *time.Time     `json:"completed_at"`
	CategoryID  *uint64        `json:"category_id"`
}
