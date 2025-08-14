package service

import (
	"ringhover-go/internal/daoerrors"
	"ringhover-go/internal/domain"
	"ringhover-go/internal/domain/models"
	"ringhover-go/internal/domain/req"
	"ringhover-go/internal/domain/resp"
	"strings"
	"time"
)

// Implémentation concrète du domain.TaskService
type Service struct {
	dao domain.InterfaceDao
}

func NewModelisationService(dao domain.InterfaceDao) *Service {
	return &Service{dao: dao}
}

// GetSubTasks requests the server to get all subtasks of a task given task id
func (s *Service) GetSubTasks(taskID uint64) (subtasksList []models.Task, err error) {

	var taskExist bool
	if taskExist, err = s.dao.ExistsTask(taskID); err != nil {
		return
	}

	if !taskExist {
		return nil, daoerrors.ErrNotFound
	}

	subtasksList, err = s.dao.GetSubTaskTree(taskID)
	if err != nil {
		return
	}

	return
}

// GetRootTasks requests the server to get all root tasks with their categories.
func (s *Service) GetRootTasks() (tasksList resp.TaskList, err error) {
	var rootTasks []models.TaskWithCategory
	if rootTasks, err = s.dao.GetRootTasksWithCategories(); err != nil {
		return
	}

	tasksList = make(resp.TaskList, 0, len(rootTasks))
	for _, r := range rootTasks {
		tasksList = append(tasksList, resp.Task{
			Id:           r.Id,
			Title:        r.Title,
			Description:  r.Description,
			Status:       r.Status,
			Priority:     r.Priority,
			DueDate:      r.DueDate,
			CompletedAt:  r.CompletedAt,
			ParentTaskID: r.ParentTaskID,
			CategoryID:   r.CategoryID,
			CreatedAt:    r.CreatedAt,
			UpdatedAt:    r.UpdatedAt,
			Category: &resp.Category{
				ID:   *r.CategoryID,
				Name: r.Category.Name,
			},
		})
	}

	return
}

// CreateTask requests the server to create a new task or subtask.
func (s *Service) CreateTask(requestTask req.CreateTaskRequest) (taskCreated models.Task, err error) {
	var status string
	if requestTask.Status == nil || strings.TrimSpace(*requestTask.Status) == "" {
		status = string(models.StatusTodo)
	} else {
		status = strings.ToLower(strings.TrimSpace(*requestTask.Status))
		if !allowedStatus(status) {
			return models.Task{}, daoerrors.ErrBadInput
		}
	}

	var priority int
	if requestTask.Priority == nil {
		priority = 0
	} else {
		priority = *requestTask.Priority
		if priority < 0 || priority > 3 {
			return models.Task{}, daoerrors.ErrBadInput
		}
	}

	if requestTask.ParentTaskID != nil {
		exists, err := s.dao.ExistsTask(*requestTask.ParentTaskID)
		if err != nil {
			return models.Task{}, err
		}
		if !exists {
			return models.Task{}, daoerrors.ErrNotFound
		}
	}

	taskCreated = models.Task{
		Title:        requestTask.Title,
		Description:  requestTask.Description,
		Status:       models.Status(status),
		Priority:     priority,
		DueDate:      requestTask.DueDate,
		CategoryID:   requestTask.CategoryID,
		ParentTaskID: requestTask.ParentTaskID,
	}

	var createdId uint64
	if createdId, err = s.dao.CreateTask(taskCreated); err != nil {
		return
	}

	taskCreated.Id = createdId

	return
}

// DeleteTask requests the server to create a new task or subtask.
func (s *Service) DeleteTask(taskId uint64) (err error) {
	var taskExist bool
	if taskExist, err = s.dao.ExistsTask(taskId); err != nil {
		return
	}

	if !taskExist {
		return daoerrors.ErrNotFound
	}

	if err = s.dao.DeleteTask(taskId); err != nil {
		return
	}

	return
}

// UpdateTask requests the server to update a task.
func (s *Service) UpdateTask(taskId uint64, requestTask req.UpdateTaskRequest) (taskUpdated models.Task, err error) {
	var taskToUpdate models.Task
	var baseTask models.Task
	if baseTask, err = s.dao.GetTaskByID(taskId); err != nil {
		return
	}

	taskToUpdate = baseTask

	if requestTask.Status != nil {
		sv := strings.ToLower(strings.TrimSpace(*requestTask.Status))
		if !allowedStatus(sv) {
			return models.Task{}, daoerrors.ErrBadInput
		}
		taskToUpdate.Status = models.Status(sv)
		if taskToUpdate.Status == models.StatusDone {
			now := time.Now()
			taskToUpdate.CompletedAt = &now
		}
		if baseTask.Status == models.StatusDone &&  taskToUpdate.Status != baseTask.Status {
			taskToUpdate.CompletedAt = nil
		}
	}

	if requestTask.Title != nil {
		title := strings.TrimSpace(*requestTask.Title)
		if title == "" || len(title) > 255 {
			return models.Task{}, daoerrors.ErrBadInput
		}
		taskToUpdate.Title = title
	}

	if requestTask.Priority != nil {
		if *requestTask.Priority < 0 || *requestTask.Priority > 3 {
			return models.Task{}, daoerrors.ErrBadInput
		}
		taskToUpdate.Priority = *requestTask.Priority
	}

	if requestTask.DueDate != nil {
		taskToUpdate.DueDate = requestTask.DueDate
	}
	if requestTask.Description != nil {
		taskToUpdate.Description = requestTask.Description
	}
	if requestTask.CategoryID != nil {
		taskToUpdate.CategoryID = requestTask.CategoryID
	}

	if requestTask.ParentTaskID != nil {
		if *requestTask.ParentTaskID == taskId {
			return models.Task{}, daoerrors.ErrBadInput
		}
		taskToUpdate.ParentTaskID = requestTask.ParentTaskID
	}

	taskToUpdate.Id = taskId
	if err = s.dao.UpdateTask(taskToUpdate); err != nil {
		return
	}
	taskUpdated = taskToUpdate

	return
}

// allowedStatus checks if the given status is allowed.
func allowedStatus(s string) bool {
	return s == string(models.StatusTodo) || s == string(models.StatusInProgress) || s == string(models.StatusDone)
}
