package repository

import (
	"database/sql"
	"github.com/Askalag/piece16/src/log"
	"github.com/Askalag/piece16/src/model"
	"github.com/jmoiron/sqlx"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func (r *TaskPostgres) GetById(id int) (*model.Task, error) {
	var task model.Task
	query := "SELECT * FROM t1.task WHERE id=$1"

	err := r.db.Get(&task, query, id)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &task, nil
	}
}

func (r *TaskPostgres) Create(obj model.Task) (int, error) {
	query := "INSERT INTO t1.task(title, tree_level) VALUES($1, $2) RETURNING id"
	row := r.db.QueryRow(query, obj.Title, obj.TreeLevel)

	var id int
	if err := row.Scan(&id); err != nil {
		log.WarnWithCode(3001, err.Error())
		return 0, err
	}
	return id, nil
}

func (r *TaskPostgres) GetAll() (*[]model.Task, error) {
	var arr []model.Task
	query := "SELECT * FROM t1.task"
	err := r.db.Select(&arr, query)
	if err != nil {
		return nil, err
	}
	return &arr, nil
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}
