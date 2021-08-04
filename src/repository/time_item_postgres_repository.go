package repository

import (
	"database/sql"
	"errors"
	"github.com/Askalag/piece16/src/log"
	"github.com/Askalag/piece16/src/model"
	"github.com/jmoiron/sqlx"
)

type TimeItemPostgres struct {
	db *sqlx.DB
}

func (r *TimeItemPostgres) Create(m *model.TimeItem) (int, error) {
	m.TreeLevel = 2
	query := "INSERT INTO t1.time_item(title, description, time_cost, tree_level, parent_id) VALUES($1, $2, $3, $4, $5) RETURNING id"
	row := r.db.QueryRow(query, m.Title, m.Description, m.TimeCost, m.TreeLevel, m.ParentId)

	var id int
	if err := row.Scan(&id); err != nil {
		log.WarnWithCode(3001, err.Error())
		return 0, err
	}
	return id, nil
}

func (r *TimeItemPostgres) GetAll() (*[]model.TimeItem, error) {
	var arr []model.TimeItem
	query := "SELECT * FROM t1.time_item"
	err := r.db.Select(&arr, query)
	if err != nil {
		return nil, err
	}
	return &arr, nil
}

func (r *TimeItemPostgres) GetById(id int) (*model.TimeItem, error) {
	var task model.TimeItem
	query := "SELECT * FROM t1.time_item WHERE id=$1"

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

func (r *TimeItemPostgres) GetByParentId(parentId int) (*[]model.TimeItem, error) {
	var task []model.TimeItem
	query := "SELECT * FROM t1.time_item WHERE parent_id=$1"

	err := r.db.Get(&task, query, parentId)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &task, nil
	}
}

func (r *TimeItemPostgres) Update(m *model.TimeItem) error {
	if m == nil || m.Id <= 0 {
		log.WarnWithCode(3003)
		return errors.New("bad params for update")
	}

	query := "UPDATE t1.task_item SET title=$1, parent_id=$2 WHERE id=$3"
	_, err := r.db.Query(query, m.Title, m.ParentId, m.Id)
	return err
}

func (r *TimeItemPostgres) DeleteById(id int) error {
	query := "DELETE FROM t1.time_item WHERE id=$1"
	_, err := r.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *TimeItemPostgres) DeleteByParentId(id int) error {
	query := "DELETE FROM t1.time_item WHERE parent_id=$1"
	_, err := r.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func NewTimeItemPostgres(db *sqlx.DB) *TimeItemPostgres {
	return &TimeItemPostgres{db: db}
}
