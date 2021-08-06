package service

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/repository"
)

type Tree interface {
	BuildById(id int) (*model.Tree, error)
	GetById(id int) (*model.Tree, error)
	GetAll() (*[]model.Tree, error)
	DeleteById(id int) error
	UpdParentTI(ti *model.TaskItem) error
	UpdParentTTI(tii *model.TimeItem) error
	DelTI(ti *model.TimeItem) error
	DelTTI(tti *model.TimeItem) error
}

type Task interface {
	Create(task *model.Task) (int, error)
	GetAll() (*[]model.Task, error)
	GetById(id int) (*model.Task, error)
	Update(m *model.Task) error
	DeleteById(id int) error
}

type TaskItem interface {
	Create(task *model.TaskItem) (int, error)
	GetAll() (*[]model.TaskItem, error)
	GetById(id int) (*model.TaskItem, error)
	GetByParentId(id int) (*model.TaskItem, error)
	Update(m *model.TaskItem) error
	UpdateParentId(old int, new int) error
	DeleteById(id int) error
	DeleteByParentId(id int) error
}

type TaskTimeItem interface {
	Create(m *model.TimeItem) (int, error)
	GetAll() (*[]model.TimeItem, error)
	GetById(id int) (*model.TimeItem, error)
	GetByParentId(id int) (*[]model.TimeItem, error)
	Update(m *model.TimeItem) error
	UpdateParentId(old int, new int) error
	DeleteById(id int) error
	DeleteByParentId(id int) error
}

type Service struct {
	TreeService         Tree
	TaskService         Task
	TaskItemService     TaskItem
	TaskTimeItemService TaskTimeItem
}

func NewService(r *repository.Repo) *Service {
	return &Service{
		TreeService:         NewTreeService(r),
		TaskService:         NewTaskService(r.T),
		TaskItemService:     NewTaskItemService(r.TI),
		TaskTimeItemService: NewTaskTimeItemService(r.TTI),
	}
}
