package repository

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	C   CmdRepo
	TR  TreeRepo
	T   TaskRepo
	TI  TaskItemRepo
	TTI TaskTimeItemRepo
}

type CmdRepo interface {
	DropAll() error
	InitTables() error
}

type TreeRepo interface {
	GetAll() (*[]model.Tree, error)
	Create(m *model.Tree) (int, error)
	Update(m *model.Tree) error
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
	DeleteByParentId(id int) error
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
		C:   postgres.NewCmdPostgres(db),
		TR:  postgres.NewTreePostgres(db),
		T:   postgres.NewTaskPostgres(db),
		TI:  postgres.NewTaskItemPostgres(db),
		TTI: postgres.NewTimeItemPostgres(db),
	}
}
