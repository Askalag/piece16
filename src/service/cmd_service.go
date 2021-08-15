package service

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/repository"
)

type CmdService struct {
	repo *repository.Repo
}

// InitTables init tables and schemas
func (c CmdService) InitTables() error {
	return c.repo.C.InitTables()
}

// DropAll drop all tables
func (c CmdService) DropAll() error {
	return c.repo.C.DropAll()
}

// FillFullTree filling tree with inner elements
func (c CmdService) FillFullTree() error {

	// filling tree
	tree := &model.Tree{Title: "tree title 1"}
	treeId, err := c.repo.TR.Create(tree)
	if err != nil {
		return err
	}

	// filling []model.Task
	var tasks = []model.Task{
		{Title: "task title 1", TreeId: treeId},
		{Title: "task title 2", TreeId: treeId},
		{Title: "task title 3", TreeId: treeId},
		{Title: "task title 4", TreeId: treeId},
		{Title: "task title 5", TreeId: treeId},
		{Title: "task title 6", TreeId: treeId},
		{Title: "task title 7", TreeId: treeId},
	}
	var tasksIds []int // count 7 - (1)
	for i := range tasks {
		id, err := c.repo.T.Create(&tasks[i])
		if err != nil {
			return err
		}
		tasksIds = append(tasksIds, id)
	}

	// filling []model.TaskItem
	var taskItems = []model.TaskItem{
		{Title: "taskItem title 1", ParentId: tasksIds[0]},
		{Title: "taskItem title 2", ParentId: tasksIds[0]},
		{Title: "taskItem title 3", ParentId: tasksIds[0]},
		{Title: "taskItem title 4", ParentId: tasksIds[1]},
		{Title: "taskItem title 5", ParentId: tasksIds[1]},
		{Title: "taskItem title 6", ParentId: tasksIds[2]},
		{Title: "taskItem title 7", ParentId: tasksIds[3]},
		{Title: "taskItem title 8", ParentId: tasksIds[4]},
		{Title: "taskItem title 9", ParentId: tasksIds[4]},
		{Title: "taskItem title 10", ParentId: tasksIds[4]},
		{Title: "taskItem title 11", ParentId: tasksIds[4]},
	}
	var taskItemsIds []int // count 11 - (1)
	for i := range taskItems {
		id, err := c.repo.TI.Create(&taskItems[i])
		if err != nil {
			return err
		}
		taskItemsIds = append(taskItemsIds, id)
	}

	// filling []model.TimeItem
	var timeItems = []model.TimeItem{
		{Title: "TimeItem title 1", Description: "timeItem Description 1", TimeCost: 12, ParentId: taskItemsIds[0]},
		{Title: "TimeItem title 2", Description: "timeItem Description 2", TimeCost: 2, ParentId: taskItemsIds[0]},
		{Title: "TimeItem title 3", Description: "timeItem Description 3", TimeCost: 123, ParentId: taskItemsIds[1]},
		{Title: "TimeItem title 4", Description: "timeItem Description 4", TimeCost: 76, ParentId: taskItemsIds[1]},
		{Title: "TimeItem title 5", Description: "timeItem Description 5", TimeCost: 3, ParentId: taskItemsIds[2]},
		{Title: "TimeItem title 6", Description: "timeItem Description 6", TimeCost: 8, ParentId: taskItemsIds[3]},
		{Title: "TimeItem title 7", Description: "timeItem Description 7", TimeCost: 3, ParentId: taskItemsIds[4]},
		{Title: "TimeItem title 8", Description: "timeItem Description 8", TimeCost: 7, ParentId: taskItemsIds[5]},
		{Title: "TimeItem title 9", Description: "timeItem Description 9", TimeCost: 18, ParentId: taskItemsIds[5]},
		{Title: "TimeItem title 10", Description: "timeItem Description 10", TimeCost: 41, ParentId: taskItemsIds[5]},
		{Title: "TimeItem title 11", Description: "timeItem Description 11", TimeCost: 13, ParentId: taskItemsIds[5]},
		{Title: "TimeItem title 12", Description: "timeItem Description 12", TimeCost: 2, ParentId: taskItemsIds[5]},
		{Title: "TimeItem title 13", Description: "timeItem Description 13", TimeCost: 1, ParentId: taskItemsIds[5]},
		{Title: "TimeItem title 14", Description: "timeItem Description 14", TimeCost: 1, ParentId: taskItemsIds[6]},
		{Title: "TimeItem title 15", Description: "timeItem Description 15", TimeCost: 1, ParentId: taskItemsIds[7]},
		{Title: "TimeItem title 16", Description: "timeItem Description 16", TimeCost: 7, ParentId: taskItemsIds[8]},
		{Title: "TimeItem title 17", Description: "timeItem Description 17", TimeCost: 4, ParentId: taskItemsIds[8]},
		{Title: "TimeItem title 18", Description: "timeItem Description 18", TimeCost: 3, ParentId: taskItemsIds[8]},
		{Title: "TimeItem title 19", Description: "timeItem Description 19", TimeCost: 1, ParentId: taskItemsIds[9]},
		{Title: "TimeItem title 20", Description: "timeItem Description 20", TimeCost: 2, ParentId: taskItemsIds[9]},
		{Title: "TimeItem title 21", Description: "timeItem Description 21", TimeCost: 12, ParentId: taskItemsIds[10]},
	}
	for i := range timeItems {
		if _, err := c.repo.TTI.Create(&timeItems[i]); err != nil {
			return err
		}
	}
	return nil
}

func NewCmdService(r *repository.Repo) *CmdService {
	return &CmdService{repo: r}
}
