package resp

import (
	"ringhover-go/internal/domain/models"
	"time"
)

type Task struct {
	ID           uint64        `json:"id"`
	Title        string        `json:"title"`
	Description  *string       `json:"description,omitempty"`
	Status       models.Status `json:"status"`
	Priority     int           `json:"priority"`
	DueDate      *time.Time    `json:"due_date,omitempty"`
	CompletedAt  *time.Time    `json:"completed_at,omitempty"`
	ParentTaskID *uint64       `json:"parent_task_id,omitempty"`
	CategoryID   *uint64       `json:"category_id,omitempty"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	Category     *Category     `json:"category,omitempty"`
}

// TaskTree: pour GET /tasks/:id/subtasks
type TaskTree struct {
	Task
	Children []TaskTree `json:"children"`
}
