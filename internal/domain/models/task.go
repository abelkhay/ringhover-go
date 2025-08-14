package models

import (
	"time"
)

// Task table metadata.
type Task struct {
	Id           uint64     `db:"id"`
	Title        string     `db:"title"`
	Description  *string    `db:"description"`
	Status       Status     `db:"status"`
	Priority     int        `db:"priority"`
	DueDate      *time.Time `db:"due_date"`
	CompletedAt  *time.Time `db:"completed_at"`
	ParentTaskID *uint64    `db:"parent_task_id"`
	CategoryID   *uint64    `db:"category_id"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at"`
}

type TaskWithCategory struct {
	Task
	Category Category `db:"category"`
}
