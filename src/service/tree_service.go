package service

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/repository"
	"sync"
)

type TreeService struct {
	trRepo  repository.TreeRepo
	tRepo   repository.TaskRepo
	tIRepo  repository.TaskItemRepo
	tTIRepo repository.TaskTimeItemRepo
}

// DeleteFullTree delete Tree include inner elements if exists.
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

// BuildById Build a new Tree by id if exists with inner elements.
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
					taskItemIds = append(taskItemIds, item.Id)
				}

				timeItems, err := t.tTIRepo.GetByParentIds(taskItemIds)
				if err != nil {
					return &res, err
				}
				times = *timeItems
			}
		}
		res = fillAndCalcTree(*tree, tasks, items, times)
		return &res, err

	}
	return nil, err
}

// GetById receive model.Tree by Id
func (t TreeService) GetById(id int) (*model.Tree, error) {
	return t.trRepo.GetById(id)
}

// GetAll receive a list of model.Tree
func (t TreeService) GetAll() (*[]model.Tree, error) {
	return t.trRepo.GetAll()
}

// DeleteById delete model.Tree by Id
func (t TreeService) DeleteById(id int) error {
	return t.trRepo.DeleteById(id)
}

// UpdTI update model.TaskItem and rebuild tree
func (t TreeService) UpdTI(treeId int, ti *model.TaskItem) (*model.Tree, error) {
	if err := t.tIRepo.Update(ti); err != nil {
		return nil, err
	}
	return t.BuildById(treeId)
}

// UpdTTI update model.TimeItem and rebuild tree
func (t TreeService) UpdTTI(treeId int, tii *model.TimeItem) (*model.Tree, error) {
	if err := t.tTIRepo.Update(tii); err != nil {
		return nil, err
	}
	return t.BuildById(treeId)
}

// DelTI delete model.TaskItem and rebuild tree
func (t TreeService) DelTI(treeId int, tiId int, deep bool) (*model.Tree, error) {
	if err := t.tIRepo.DeleteById(tiId); err != nil {
		return nil, err
	}
	return t.BuildById(treeId)
}

// DelTTI delete model.TimeItem and rebuild tree
func (t TreeService) DelTTI(treeId int, ttiId int, deep bool) (*model.Tree, error) {
	if err := t.tTIRepo.DeleteById(ttiId); err != nil {
		return nil, err
	}
	return t.BuildById(treeId)
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

// calcTask calculate model.Task TimeCostAverage and TimeCostTotal
func calcTask(chCalc chan model.Task, t model.Task, items []model.TaskItem) {
	for _, e := range items {
		t.TimeCostTotal += e.TimeCostTotal
		t.TimeCostAverage++
	}
	if t.TimeCostAverage != 0 {
		t.TimeCostAverage = t.TimeCostTotal / t.TimeCostAverage
	}
	chCalc <- t
}

// calcTaskItem calculate model.TaskItem TimeCostAverage and TimeCostTotal
func calcTaskItem(chCalc chan model.TaskItem, ti model.TaskItem, timeItems []model.TimeItem) {
	for _, e := range timeItems {
		ti.TimeCostTotal += e.TimeCost
		ti.TimeCostAverage++
	}
	if ti.TimeCostAverage > 0 {
		ti.TimeCostAverage = ti.TimeCostTotal / ti.TimeCostAverage
	}
	chCalc <- ti
}

// fillTree filling full tree include inner elements.
func fillAndCalcTree(tree model.Tree, tasks []model.Task, taskItems []model.TaskItem, times []model.TimeItem) model.Tree {
	var wg sync.WaitGroup
	chTICalc := make(chan model.TaskItem)
	chTCalc := make(chan model.Task)

	// filling times to taskItems
	go func() {
		wg.Add(1)
		for i, e := range taskItems {
			for _, e2 := range times {
				if e.Id == e2.ParentId {
					taskItems[i].TimeItems = append(taskItems[i].TimeItems, e2)
				}
			}
			go calcTaskItem(chTICalc, taskItems[i], taskItems[i].TimeItems)
			taskItems[i] = <-chTICalc
		}
		wg.Done()
	}()

	wg.Wait()

	// filling taskItems to tasks
	go func() {
		wg.Add(1)
		for i, e := range tasks {
			for _, e2 := range taskItems {
				if e.Id == e2.ParentId {
					tasks[i].TaskItems = append(tasks[i].TaskItems, e2)
				}
			}
			go calcTask(chTCalc, tasks[i], tasks[i].TaskItems)
			tasks[i] = <-chTCalc
		}
		wg.Done()
	}()

	wg.Wait()

	tree.Tasks = tasks
	return tree
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
