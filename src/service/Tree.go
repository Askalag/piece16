package service

import (
	"fmt"
	"github.com/Askalag/piece16/src/repository"
)

type Tree struct {
	task repository.TaskRepo
	taskItem repository.TaskItemRepo
	timeItem repository.TimeItemRepo
}

func (t *Tree) Create() {
	fmt.Println(".end")

}

func NewTreeService(r *repository.Repo) *Tree {
	return &Tree{
		task: r.TaskRepo,
		taskItem: r.TaskItemRepo,
		timeItem: r.TimeItemRepo,
	}
}
