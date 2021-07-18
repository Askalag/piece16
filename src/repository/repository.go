package repository

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	TaskRepo
	TaskItemRepo
	TimeItemRepo
}

type TaskRepo interface {
	Create(task model.Task) (int, error)
	GetById(id int) (*model.Task, error)
	GetAll() (*[]model.Task, error)
}

type TaskItemRepo interface {
}

type TimeItemRepo interface {
}

func NewTreeRepository(db *sqlx.DB) *Repo {
	return &Repo{
		TaskRepo:     NewTaskPostgres(db),
		TaskItemRepo: NewTaskItemPostgres(db),
		TimeItemRepo: NewTimeItemPostgres(db),
	}
}
