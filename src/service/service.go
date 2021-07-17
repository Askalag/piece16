package service

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/repository"
)

type Task interface {
	Create(task model.Task) (int64, error)
	GetAll() ([]model.Task, error)
}

type TaskItem interface {
	Create(task model.TaskItem) (int64, error)
	GetAll() ([]model.TaskItem, error)
}

type TaskTimeItem interface {
}

type Service struct {
	Task
	TaskItem
	TaskTimeItem
}

func NewService(r *repository.Repo) *Service {
	return &Service{
		Task:     NewTaskService(r),
		TaskItem: NewTaskItemService(r),
	}
}
