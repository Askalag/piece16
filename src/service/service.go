package service

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/repository"
)

type Task interface {
	Create(task *model.Task) (int, error)
	GetAll() (*[]model.Task, error)
	GetById(id int) (*model.Task, error)
	Update(m *model.Task) error
	DeleteById(id int) error
}

type TaskItem interface {
	Create(task model.TaskItem) (int64, error)
	GetAll() ([]model.TaskItem, error)
}

type TaskTimeItem interface {
}

type Service struct {
	TaskService         Task
	TaskItemService     TaskItem
	TaskTimeItemService TaskTimeItem
}

func NewService(r *repository.Repo) *Service {
	return &Service{
		TaskService:     NewTaskService(r),
		TaskItemService: NewTaskItemService(r),
	}
}
