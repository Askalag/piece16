package service

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/repository"
)

type TreeService struct {
	tRepo   repository.TaskRepo
	tIRepo  repository.TaskItemRepo
	tTIRepo repository.TaskTimeItemRepo
}

func (t TreeService) BuildById(id int) (*model.Tree, error) {
	panic("implement me")
}

func (t TreeService) GetById(id int) (*model.Tree, error) {
	panic("implement me")
}

func (t TreeService) GetAll() (*[]model.Tree, error) {
	panic("implement me")
}

func (t TreeService) DeleteById(id int) error {
	panic("implement me")
}

func (t TreeService) UpdParentTI(ti *model.TaskItem) error {
	panic("implement me")
}

func (t TreeService) UpdParentTTI(tii *model.TimeItem) error {
	panic("implement me")
}

func (t TreeService) DelTI(ti *model.TimeItem) error {
	panic("implement me")
}

func (t TreeService) DelTTI(tti *model.TimeItem) error {
	panic("implement me")
}

func NewTreeService(r *repository.Repo) *TreeService {
	return &TreeService{
		tRepo:   r.T,
		tIRepo:  r.TI,
		tTIRepo: r.TTI,
	}
}
