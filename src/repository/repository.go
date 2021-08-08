package repository

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	TR  TreeRepo
	T   TaskRepo
	TI  TaskItemRepo
	TTI TaskTimeItemRepo
}

type TreeRepo interface {
	Create(m *model.Tree) (int, error)
	GetAll() (*[]model.Tree, error)
	GetById(id int) (*model.Tree, error)
	DeleteById(id int) error
}

type TaskRepo interface {
	Create(m *model.Task) (int, error)
	GetById(id int) (*model.Task, error)
	GetByIds(ids []int) (*[]model.Task, error)
	GetByTreeId(id int) (*[]model.Task, error)
	GetAll() (*[]model.Task, error)
	Update(m *model.Task) error
	DeleteById(id int) error
	DeleteByIds(ids []int) error
	DeleteByTreeId(id int) error
}

type TaskItemRepo interface {
	Create(m *model.TaskItem) (int, error)
	GetAll() (*[]model.TaskItem, error)
	GetById(id int) (*model.TaskItem, error)
	GetByIds(ids []int) (*[]model.TaskItem, error)
	GetByParentId(id int) (*[]model.TaskItem, error)
	GetByParentIds(ids []int) (*[]model.TaskItem, error)
	Update(m *model.TaskItem) error
	DeleteById(id int) error
	DeleteByIds(ids []int) error
}

type TaskTimeItemRepo interface {
	Create(m *model.TimeItem) (int, error)
	GetAll() (*[]model.TimeItem, error)
	GetById(id int) (*model.TimeItem, error)
	GetByIds(ids []int) (*[]model.TimeItem, error)
	GetByParentId(id int) (*[]model.TimeItem, error)
	GetByParentIds(ids []int) (*[]model.TimeItem, error)
	Update(m *model.TimeItem) error
	DeleteById(id int) error
	DeleteByIds(id []int) error
}

func NewTreeRepository(db *sqlx.DB) *Repo {
	return &Repo{
		TR:  NewTreePostgres(db),
		T:   NewTaskPostgres(db),
		TI:  NewTaskItemPostgres(db),
		TTI: NewTimeItemPostgres(db),
	}
}
