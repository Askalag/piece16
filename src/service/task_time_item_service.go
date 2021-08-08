package service

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/repository"
)

type TaskTimeItemService struct {
	repo repository.TaskTimeItemRepo
}

func (s *TaskTimeItemService) DeleteByIds(ids []int) error {
	return s.repo.DeleteByIds(ids)
}

func (s *TaskTimeItemService) Create(m *model.TimeItem) (int, error) {
	return s.repo.Create(m)
}

func (s *TaskTimeItemService) GetById(id int) (*model.TimeItem, error) {
	return s.repo.GetById(id)
}

func (s *TaskTimeItemService) GetByParentId(id int) (*[]model.TimeItem, error) {
	return s.repo.GetByParentId(id)
}

func (s *TaskTimeItemService) GetAll() (*[]model.TimeItem, error) {
	return s.repo.GetAll()
}

func (s *TaskTimeItemService) Update(m *model.TimeItem) error {
	return s.repo.Update(m)
}

func (s *TaskTimeItemService) UpdateParentId(old int, new int) error {
	// toDo
	return nil
}

func (s *TaskTimeItemService) DeleteById(id int) error {
	return s.repo.DeleteById(id)
}

func (s *TaskTimeItemService) DeleteByParentId(id int) error {
	// toDo
	return nil
}

func NewTaskTimeItemService(r repository.TaskTimeItemRepo) *TaskTimeItemService {
	return &TaskTimeItemService{repo: r}
}
