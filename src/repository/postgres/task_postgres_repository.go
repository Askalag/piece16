package postgres

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

func (r *TaskPostgres) GetByTreeId(id int) (*[]model.Task, error) {
	var res []model.Task
	query := "SELECT * FROM t1.task WHERE tree_id=$1"
	err := r.db.Select(&res, query, id)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

func (r *TaskPostgres) GetByIds(ids []int) (*[]model.Task, error) {
	var res []model.Task
	rawQuery := "SELECT * FROM t1.task WHERE id IN (?)"
	query, args, err := sqlx.In(rawQuery, ids)
	if err != nil {
		return &res, err
	}

	query = r.db.Rebind(query)

	err = r.db.Select(&res, query, args...)
	return &res, err
}

func (r *TaskPostgres) DeleteByIds(ids []int) error {
	rawQuery := "DELETE FROM t1.task WHERE id IN (?)"
	query, args, err := sqlx.In(rawQuery, ids)
	if err != nil {
		return err
	}

	query = r.db.Rebind(query)

	_, err = r.db.Query(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskPostgres) DeleteByTreeId(id int) error {
	query := "DELETE FROM t1.task WHERE tree_id=$1"
	_, err := r.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskPostgres) GetById(id int) (*model.Task, error) {
	var res model.Task
	query := "SELECT * FROM t1.task WHERE id=$1"

	err := r.db.Get(&res, query, id)

	if err != nil {
		log.WarnWithCode(3001, err.Error())
		switch {
		case err == sql.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}
	}
	return &res, nil
}

func (r *TaskPostgres) Create(m *model.Task) (int, error) {
	m.TreeLevel = 1
	query := "INSERT INTO t1.task(title, tree_level, tree_id) VALUES($1, $2, $3) RETURNING id"
	row := r.db.QueryRow(query, m.Title, m.TreeLevel, m.TreeId)

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
	m.TreeLevel = 1
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
