package domain

import "ringhover-go/internal/domain/models"

type ModelisationServiceInterface interface {
	GetSubTasks(taskID uint64) ([]models.Task, error)
}

type InterfaceDao interface {
	GetSubTaskTree(taskID uint64) ([]models.Task, error)
	ExistsTask(taskID uint64) (bool, error)
}
