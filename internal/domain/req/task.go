package req

import (
	"time"
)

// CreateTaskRequest request struct onto which client requests should be successfully mapped.
type CreateTaskRequest struct {
	Title        string     `json:"title"`
	Description  *string    `json:"description,omitempty"`
	Status       *string    `json:"status,omitempty"`   // todo
	Priority     *int       `json:"priority,omitempty"` // default: 0
	DueDate      *time.Time `json:"due_date,omitempty"`
	CategoryID   *uint64    `json:"category_id,omitempty"`
	ParentTaskID *uint64    `json:"parent_task_id,omitempty"` // null => root task
}

// UpdateTaskRequest request struct onto which client requests should be successfully mapped.
type UpdateTaskRequest struct {
	Title        *string    `json:"title"`
	Description  *string    `json:"description"`
	Status       *string    `json:"status"`
	Priority     *int       `json:"priority"`
	DueDate      *time.Time `json:"due_date"`
	CategoryID   *uint64    `json:"category_id,omitempty"`
	ParentTaskID *uint64    `json:"parent_task_id,omitempty"`
}
