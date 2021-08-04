package repository

import (
	"database/sql"
	"errors"
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

func (r *TaskPostgres) Create(m *model.Task) (int, error) {
	m.TreeLevel = 0
	query := "INSERT INTO t1.task(title, tree_level) VALUES($1, $2) RETURNING id"
	row := r.db.QueryRow(query, m.Title, m.TreeLevel)

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

func (r *TaskPostgres) Update(m *model.Task) error {
	if m == nil || m.Id <= 0 {
		log.WarnWithCode(3003)
		return errors.New("bad params for update")
	}

	query := "UPDATE t1.task SET title=$1 WHERE id=$2"
	_, err := r.db.Query(query, m.Title, m.Id)
	return err
}

func (r *TaskPostgres) DeleteById(id int) error {
	query := "DELETE FROM t1.task WHERE id=$1"
	_, err := r.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}
