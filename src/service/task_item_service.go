package service

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/repository"
)

type TaskItemService struct {
	repo repository.TaskItemRepo
}

func (s TaskItemService) Create(task model.TaskItem) (int64, error) {
	panic("implement me")
}

func (s TaskItemService) GetAll() ([]model.TaskItem, error) {
	panic("implement me")
}

func NewTaskItemService(r repository.TaskItemRepo) *TaskItemService {
	return &TaskItemService{repo: r}
}
