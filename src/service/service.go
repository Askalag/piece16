package service

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/repository"
)

type Cmd interface {
	FillFullTree() error
	InitTables() error
	DropAll() error
}

type Tree interface {
	Create(body *model.Tree) (int, error)
	Update(m *model.Tree) error
	BuildById(id int) (*model.Tree, error)
	GetById(id int) (*model.Tree, error)
	GetAll() (*[]model.Tree, error)
	DeleteById(id int) error
	DeleteFullTree(id *model.Tree) error
	UpdT(treeId int, task *model.Task) (*model.Tree, error)
	UpdTI(treeId int, ti *model.TaskItem) (*model.Tree, error)
	UpdTTI(treeId int, tii *model.TimeItem) (*model.Tree, error)
	DelTI(treeId int, tiId int, deep bool) (*model.Tree, error)
	DelTTI(treeId int, ttiId int, deep bool) (*model.Tree, error)
}

type Task interface {
	Create(task *model.Task) (int, error)
	GetAll() (*[]model.Task, error)
	GetById(id int) (*model.Task, error)
	Update(m *model.Task) error
	DeleteById(id int) error
	DeleteByIds(ids []int) error
	DeleteByTreeId(id int) error
}

type TaskItem interface {
	Create(task *model.TaskItem) (int, error)
	GetAll() (*[]model.TaskItem, error)
	GetById(id int) (*model.TaskItem, error)
	GetByParentId(id int) (*[]model.TaskItem, error)
	GetByParentIds(id []int) (*[]model.TaskItem, error)
	Update(m *model.TaskItem) error
	UpdateParentId(old int, new int) error
	DeleteById(id int) error
	DeleteByIds(ids []int) error
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
	DeleteByIds(ids []int) error
	DeleteByParentId(id int) error
}

type Service struct {
	CmdService          Cmd
	TreeService         Tree
	TaskService         Task
	TaskItemService     TaskItem
	TaskTimeItemService TaskTimeItem
}

func NewService(r *repository.Repo) *Service {
	return &Service{
		CmdService:          NewCmdService(r),
		TreeService:         NewTreeService(r),
		TaskService:         NewTaskService(r.T),
		TaskItemService:     NewTaskItemService(r.TI),
		TaskTimeItemService: NewTaskTimeItemService(r.TTI),
	}
}
