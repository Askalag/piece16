package repository

import (
	"database/sql"
	"errors"
	"github.com/Askalag/piece16/src/log"
	"github.com/Askalag/piece16/src/model"
	"github.com/jmoiron/sqlx"
)

type TaskItemPostgres struct {
	db *sqlx.DB
}

func (r *TaskItemPostgres) GetByParentIds(ids []int) (*[]model.TaskItem, error) {
	var res []model.TaskItem
	rawQuery := "SELECT * FROM t1.task_item WHERE parent_id IN (?)"
	query, args, err := sqlx.In(rawQuery, ids)
	if err != nil {
		return &res, err
	}

	query = r.db.Rebind(query)

	err = r.db.Select(&res, query, args...)
	return &res, err
}

func (r *TaskItemPostgres) GetByIds(ids []int) (*[]model.TaskItem, error) {
	var res []model.TaskItem
	rawQuery := "SELECT * FROM t1.task_item WHERE id IN (?)"
	query, args, err := sqlx.In(rawQuery, ids)
	if err != nil {
		return &res, err
	}

	query = r.db.Rebind(query)

	err = r.db.Select(&res, query, args...)
	return &res, err
}

func (r *TaskItemPostgres) DeleteByIds(ids []int) error {
	rawQuery := "DELETE FROM t1.task_item WHERE id IN (?)"
	query, args, err := sqlx.In(rawQuery, ids)
	if err != nil {
		log.WarnWithCode(3001, err.Error())
		return err
	}

	query = r.db.Rebind(query)

	_, err = r.db.Query(query, args...)
	if err != nil {
		log.WarnWithCode(3001, err.Error())
		return err
	}
	return nil
}

func (r *TaskItemPostgres) Create(m *model.TaskItem) (int, error) {
	m.TreeLevel = 2
	query := "INSERT INTO t1.task_item(title, tree_level, parent_id) VALUES($1, $2, $3) RETURNING id"
	row := r.db.QueryRow(query, m.Title, m.TreeLevel, m.ParentId)

	var id int
	if err := row.Scan(&id); err != nil {
		log.WarnWithCode(3001, err.Error())
		return 0, err
	}
	return id, nil
}

func (r *TaskItemPostgres) GetAll() (*[]model.TaskItem, error) {
	var arr []model.TaskItem
	query := "SELECT * FROM t1.task_item"
	err := r.db.Select(&arr, query)
	if err != nil {
		return nil, err
	}
	return &arr, nil
}

func (r *TaskItemPostgres) GetById(id int) (*model.TaskItem, error) {
	var res model.TaskItem
	query := "SELECT * FROM t1.task_item WHERE id=$1"

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

func (r *TaskItemPostgres) GetByParentId(parentId int) (*[]model.TaskItem, error) {
	var res []model.TaskItem
	query := "SELECT * FROM t1.task_item WHERE parent_id=$1"

	err := r.db.Get(&res, query, parentId)

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

func (r *TaskItemPostgres) Update(m *model.TaskItem) error {
	if m == nil || m.Id <= 0 {
		log.WarnWithCode(3003)
		return errors.New("bad params for update")
	}
	m.TreeLevel = 2
	query := "UPDATE t1.task_item SET title=$1, parent_id=$2 WHERE id=$3"
	_, err := r.db.Query(query, m.Title, m.ParentId, m.Id)
	return err
}

func (r *TaskItemPostgres) DeleteById(id int) error {
	query := "DELETE FROM t1.task_item WHERE id=$1"
	_, err := r.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskItemPostgres) DeleteByParentId(id int) error {
	query := "DELETE FROM t1.task_item WHERE parent_id=$1"
	_, err := r.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func NewTaskItemPostgres(db *sqlx.DB) *TaskItemPostgres {
	return &TaskItemPostgres{db: db}
}
