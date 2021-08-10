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

func (r *TimeItemPostgres) GetByParentIds(ids []int) (*[]model.TimeItem, error) {
	var res []model.TimeItem
	rawQuery := "SELECT * FROM t1.time_item WHERE parent_id IN (?)"
	query, args, err := sqlx.In(rawQuery, ids)
	if err != nil {
		return &res, err
	}

	query = r.db.Rebind(query)

	err = r.db.Select(&res, query, args...)
	return &res, err
}

func (r *TimeItemPostgres) GetByIds(ids []int) (*[]model.TimeItem, error) {
	var res []model.TimeItem
	rawQuery := "SELECT * FROM t1.time_item WHERE id IN (?)"
	query, args, err := sqlx.In(rawQuery, ids)
	if err != nil {
		return &res, err
	}

	query = r.db.Rebind(query)

	err = r.db.Select(&res, query, args...)
	return &res, err
}

func (r *TimeItemPostgres) DeleteByIds(ids []int) error {
	rawQuery := "DELETE FROM t1.time_item WHERE id IN (?)"
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

func (r *TimeItemPostgres) Create(m *model.TimeItem) (int, error) {
	m.TreeLevel = 3
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
	var res model.TimeItem
	query := "SELECT * FROM t1.time_item WHERE id=$1"

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

func (r *TimeItemPostgres) GetByParentId(parentId int) (*[]model.TimeItem, error) {
	var res []model.TimeItem
	query := "SELECT * FROM t1.time_item WHERE parent_id=$1"

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

func (r *TimeItemPostgres) Update(m *model.TimeItem) error {
	if m == nil || m.Id <= 0 {
		log.WarnWithCode(3003)
		return errors.New("bad params for update")
	}
	m.TreeLevel = 3
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
