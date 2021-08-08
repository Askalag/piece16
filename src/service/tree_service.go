package service

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/repository"
)

type TreeService struct {
	trRepo  repository.TreeRepo
	tRepo   repository.TaskRepo
	tIRepo  repository.TaskItemRepo
	tTIRepo repository.TaskTimeItemRepo
}

func (t TreeService) DeleteFullTree(tree *model.Tree) error {
	ids := obtainTreeElemIds(tree)
	var err error = nil

	if len(ids.timeItemIds) > 0 {
		err = t.tTIRepo.DeleteByIds(ids.timeItemIds)
	}

	if len(ids.taskItemIds) > 0 {
		err = t.tIRepo.DeleteByIds(ids.timeItemIds)
	}

	if len(ids.taskIds) > 0 {
		err = t.tRepo.DeleteByIds(ids.taskIds)
	}

	err = t.trRepo.DeleteById(tree.Id)

	return err
}

func (t TreeService) BuildById(id int) (*model.Tree, error) {
	var res model.Tree
	var tasks []model.Task
	var items []model.TaskItem
	var times []model.TimeItem

	tree, err := t.trRepo.GetById(id)
	if err != nil {
		return &res, err
	}

	if tree != nil {
		res = *tree
		taskElms, err := t.tRepo.GetByTreeId(tree.Id)
		if err != nil {
			return &res, err
		}

		if len(*taskElms) > 0 {
			var taskIds []int
			for _, task := range *taskElms {
				taskIds = append(taskIds, task.Id)
			}
			tasks = *taskElms

			taskItems, err := t.tIRepo.GetByParentIds(taskIds)
			if err != nil {
				return &res, err
			}
			items = *taskItems

			if len(*taskItems) > 0 {
				var taskItemIds []int
				for _, item := range *taskItems {
					taskIds = append(taskItemIds, item.Id)
				}

				timeItems, err := t.tTIRepo.GetByParentIds(taskItemIds)
				if err != nil {
					return &res, err
				}
				times = *timeItems
			}
		}
		fillTree(tree, &tasks, &items, &times)
		return &res, err
	}
	return nil, err
}

func (t TreeService) GetById(id int) (*model.Tree, error) {
	return t.trRepo.GetById(id)
}

func (t TreeService) GetAll() (*[]model.Tree, error) {
	return t.trRepo.GetAll()
}

func (t TreeService) DeleteById(id int) error {
	return t.trRepo.DeleteById(id)
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

func fillTree(tree *model.Tree, tasks *[]model.Task, taskItems *[]model.TaskItem, times *[]model.TimeItem) *model.Tree {

	// filling times to taskItems
	for _, e := range *taskItems {
		for _, e2 := range *times {
			if e.Id == e2.ParentId {
				e.TimeItems = append(e.TimeItems, e2)
			}
		}
	}

	// filling taskItems to tasks
	for _, e := range *tasks {
		for _, e2 := range *taskItems {
			if e.Id == e2.ParentId {
				e.TaskItems = append(e.TaskItems, e2)
			}
		}
	}

	tree.Tasks = *tasks
	return tree
}

func obtainTreeElemIds(tree *model.Tree) *TreeElemIds {
	ids := &TreeElemIds{}

	// taskIds
	if len(tree.Tasks) > 0 {
		taskItems := make([]model.TaskItem, 0)
		for _, item := range tree.Tasks {
			ids.taskIds = append(ids.taskIds, item.Id)
			if len(item.TaskItems) > 0 {
				taskItems = append(taskItems, item.TaskItems...)
			}

			// taskItemIds
			if len(taskItems) > 0 {
				timeItems := make([]model.TimeItem, 0)
				for _, item := range taskItems {
					ids.taskItemIds = append(ids.taskItemIds, item.Id)
					if len(item.TimeItems) > 0 {
						timeItems = append(timeItems, item.TimeItems...)
					}

					// timeItemIds
					if len(timeItems) > 0 {
						for _, item := range timeItems {
							ids.timeItemIds = append(ids.timeItemIds, item.Id)
						}
					}
				}
			}
		}
	}
	return ids
}

type TreeElemIds struct {
	taskIds     []int
	taskItemIds []int
	timeItemIds []int
}

func NewTreeService(r *repository.Repo) *TreeService {
	return &TreeService{
		trRepo:  r.TR,
		tRepo:   r.T,
		tIRepo:  r.TI,
		tTIRepo: r.TTI,
	}
}
