package domain

import (
	"ringhover-go/internal/domain/models"
	"ringhover-go/internal/domain/req"
	"ringhover-go/internal/domain/resp"
)

type ModelisationServiceInterface interface {
	GetSubTasks(taskID uint64) (subtasksList []models.Task, err error)
	GetRootTasks() (tasksList resp.TaskList, err error)
	CreateTask(requestTask req.CreateTaskRequest) (taskCreated models.Task, err error)
	DeleteTask(taskId uint64) (err error)
	UpdateTask(taskId uint64, requestTask req.UpdateTaskRequest) (taskUpdated models.Task, err error)
}

type InterfaceDao interface {
	GetRootTasksWithCategories() (tasksList []models.TaskWithCategory, err error)
	GetSubTaskTree(taskID uint64) (subtasks []models.Task, err error)
	ExistsTask(taskID uint64) (exist bool, err error)
	GetTaskByID(taskID uint64) (task models.Task, err error)
	CreateTask(taskToCreate models.Task) (createdId uint64, err error)
	DeleteTask(taskId uint64) (err error)
	UpdateTask(taskToUpdate models.Task) (err error)
}
