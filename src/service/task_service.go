package service

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/repository"
)

type TaskService struct {
	repo repository.TaskRepo
}

func (s *TaskService) Create(task model.Task) (int64, error) {
	return 0, nil
}

func (s *TaskService) GetAll() ([]model.Task, error) {
	return nil, nil
}

func NewTaskService(r repository.TaskRepo) *TaskService {
	return &TaskService{repo: r}
}
