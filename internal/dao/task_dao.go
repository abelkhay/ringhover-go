package dao

import (
	"database/sql"
	"errors"
	"ringhover-go/internal/daoerrors"
	"ringhover-go/internal/domain/models"

	"github.com/jmoiron/sqlx"
)

type Dao struct {
	db *sqlx.DB
}

func NewDao(db *sqlx.DB) *Dao { return &Dao{db: db} }

// GetSubTaskTree returns all of subtasks according to given task ID.
func (d *Dao) GetSubTaskTree(taskID uint64) (subtasks []models.Task, err error) {
	const query = `
		WITH RECURSIVE sub AS (
			SELECT t.* FROM tasks t WHERE t.parent_task_id = ?
			UNION ALL
			SELECT t.* FROM tasks t
			JOIN sub s ON t.parent_task_id = s.id
		)
		SELECT * FROM sub
		ORDER BY parent_task_id, id;`

	if err = d.db.Select(&subtasks, query, taskID); err != nil {
		return
	}

	return
}

// ExistsTask return true if current task Id exist.
func (d *Dao) ExistsTask(taskID uint64) (exist bool, err error) {
	const query = `SELECT EXISTS(SELECT 1 FROM tasks WHERE id = ?) AS exist`
	if err := d.db.Get(&exist, query, taskID); err != nil {
		return false, err
	}
	return
}

// GetTaskByID returns a complete task with given task id.
func (d *Dao) GetTaskByID(taskID uint64) (task models.Task, err error) {
	const query = `
		SELECT
			id, title, description, status, priority, due_date, completed_at,
			parent_task_id, category_id, created_at, updated_at
		FROM tasks
		WHERE id = ?
	`
	if err := d.db.Get(&task, query, taskID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Task{}, daoerrors.ErrNotFound
		}
		return models.Task{}, err
	}
	return
}

// GetRootTasksWithCategories returns all root tasks with their categories.
func (d *Dao) GetRootTasksWithCategories() (tasksList []models.TaskWithCategory, err error) {
	const query = `
		SELECT
			t.*,
			c.id   AS "category.id",
			c.name AS "category.name"
		FROM tasks AS t
		LEFT JOIN categories AS c ON c.id = t.category_id
		WHERE t.parent_task_id IS NULL
	`
	if err = d.db.Select(&tasksList, query); err != nil {
		return
	}

	return
}

// CreateTask add a new task or subtask into the DB.
func (d *Dao) CreateTask(taskToCreate models.Task) (createdId uint64, err error) {
	const query = `
		INSERT INTO tasks (
			title, description, status, priority, due_date, parent_task_id, category_id
		) VALUES (
			:title, :description, :status, :priority, :due_date, :parent_task_id, :category_id
		)
	`
	res, err := d.db.NamedExec(query, &taskToCreate)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	createdId = uint64(id)
	return
}

// DeleteTask delete a task and all associated subtask into the DB.
func (d *Dao) DeleteTask(taskId uint64) (err error) {
	const query = `DELETE FROM tasks WHERE id = ?`
	_, err = d.db.Exec(query, taskId)
	if err != nil {
		return
	}
	return
}

// UpdateTask add a new task or subtask into the DB.
func (d *Dao) UpdateTask(taskToUpdate models.Task) (err error) {
	const query = `
	UPDATE tasks SET
			title           = :title,
			description     = :description,
			status          = :status,
			priority        = :priority,
			due_date        = :due_date,
			completed_at    = :completed_at,
			parent_task_id  = :parent_task_id,
			category_id     = :category_id
		WHERE id = :id
	`
	res, err := d.db.NamedExec(query, &taskToUpdate)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return daoerrors.ErrNotFound
	}

	return
}
