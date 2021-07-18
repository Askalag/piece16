package service

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/repository"
)

type TaskService struct {
	repo repository.TaskRepo
}

func (s *TaskService) GetById(id int) (*model.Task, error) {
	return s.repo.GetById(id)
}

func (s *TaskService) Create(task model.Task) (int, error) {
	return s.repo.Create(task)
}

func (s *TaskService) GetAll() (*[]model.Task, error) {
	return s.repo.GetAll()
}

func NewTaskService(r repository.TaskRepo) *TaskService {
	return &TaskService{repo: r}
}
