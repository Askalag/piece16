package repository

import "github.com/jmoiron/sqlx"

type TaskItemPostgres struct {
	db *sqlx.DB
}

func NewTaskItemPostgres(db *sqlx.DB) *TaskItemPostgres {
	return &TaskItemPostgres{db: db}
}
