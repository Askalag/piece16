package service

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/repository"
)

type TaskItemService struct {
	repo repository.TaskItemRepo
}

func (s *TaskItemService) DeleteById(id int) error {
	return nil
}

func (s *TaskItemService) Create(task model.TaskItem) (int64, error) {
	return 0, nil
}

func (s *TaskItemService) GetAll() ([]model.TaskItem, error) {
	return nil, nil
}

func NewTaskItemService(r repository.TaskItemRepo) *TaskItemService {
	return &TaskItemService{repo: r}
}
