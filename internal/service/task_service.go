package service

import (
	"ringhover-go/internal/daoerrors"
	"ringhover-go/internal/domain"
	"ringhover-go/internal/domain/models"
)

// Implémentation concrète du domain.TaskService
type Service struct {
	dao domain.InterfaceDao
}

func NewModelisationService(dao domain.InterfaceDao) *Service {
	return &Service{dao: dao}
}

// GetSubTasks requests the server to get all subtaks of a task given task id
func (s *Service) GetSubTasks(taskID uint64) ([]models.Task, error) {

	var taskExist bool
	var err error

	if taskExist, err = s.dao.ExistsTask(taskID); err != nil {
		return nil, err
	}

	if !taskExist {
		return nil, daoerrors.ErrNotFound
	}

	rows, err := s.dao.GetSubTaskTree(taskID)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
