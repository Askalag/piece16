package handler

import "github.com/Askalag/piece16/src/service"

type TaskItemHandler struct {
	s service.TaskItem
}

func NewTaskItemHandler(s service.TaskItem) *TaskItemHandler {
	return &TaskItemHandler{s: s}
}
