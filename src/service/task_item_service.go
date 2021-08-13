package service

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/repository"
)

type TaskItemService struct {
	repo repository.TaskItemRepo
}

func (s *TaskItemService) GetByParentIds(ids []int) (*[]model.TaskItem, error) {
	return s.repo.GetByParentIds(ids)
}

func (s *TaskItemService) DeleteByIds(ids []int) error {
	return s.repo.DeleteByIds(ids)
}

func (s *TaskItemService) Create(m *model.TaskItem) (int, error) {
	return s.repo.Create(m)
}

func (s *TaskItemService) GetById(id int) (*model.TaskItem, error) {
	return s.repo.GetById(id)
}

func (s *TaskItemService) GetByParentId(id int) (*[]model.TaskItem, error) {
	return s.repo.GetByParentId(id)
}

func (s *TaskItemService) GetAll() (*[]model.TaskItem, error) {
	return s.repo.GetAll()
}

func (s *TaskItemService) Update(m *model.TaskItem) error {
	return s.repo.Update(m)
}

func (s *TaskItemService) UpdateParentId(old int, new int) error {
	// toDo
	return nil
}

func (s *TaskItemService) DeleteById(id int) error {
	return s.repo.DeleteById(id)
}

func (s *TaskItemService) DeleteByParentId(id int) error {
	return s.repo.DeleteByParentId(id)
}

func NewTaskItemService(r repository.TaskItemRepo) *TaskItemService {
	return &TaskItemService{repo: r}
}
