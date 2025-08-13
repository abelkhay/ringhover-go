package dao

import (
	"ringhover-go/internal/domain/models"

	"github.com/jmoiron/sqlx"
)

type Dao struct {
	db *sqlx.DB
}

func NewDao(db *sqlx.DB) *Dao { return &Dao{db: db} }

// GetSubTaskTree returns all of subtasks according to given task ID.
func (d *Dao) GetSubTaskTree(taskID uint64) ([]models.Task, error) {
	const query = `
		WITH RECURSIVE sub AS (
			SELECT t.* FROM tasks t WHERE t.parent_task_id = ?
			UNION ALL
			SELECT t.* FROM tasks t
			JOIN sub s ON t.parent_task_id = s.id
		)
		SELECT * FROM sub
		ORDER BY parent_task_id, id;`
	var rows []models.Task
	if err := d.db.Select(&rows, query, taskID); err != nil {
		return nil, err
	}
	return rows, nil
}

// ExistsTask return true if current task Id exist.
func (d *Dao) ExistsTask(taskID uint64) (bool, error) {
	const query = `SELECT EXISTS(SELECT 1 FROM tasks WHERE id = ?) AS ok`
	var ok bool
	if err := d.db.Get(&ok, query, taskID); err != nil {
		return false, err
	}
	return ok, nil
}
