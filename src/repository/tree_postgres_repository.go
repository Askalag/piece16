package repository

import (
	"database/sql"
	"errors"
	"github.com/Askalag/piece16/src/log"
	"github.com/Askalag/piece16/src/model"
	"github.com/jmoiron/sqlx"
)

type TreePostgres struct {
	db *sqlx.DB
}

func (r TreePostgres) Update(m *model.Tree) error {
	if m == nil || m.Id <= 0 {
		log.WarnWithCode(3003)
		return errors.New("bad params for update")
	}
	m.TreeLevel = 2
	query := "UPDATE t1.tree SET title=$1 WHERE id=$2"
	_, err := r.db.Query(query, m.Title, m.Id)
	return err
}

func (r TreePostgres) Create(m *model.Tree) (int, error) {
	m.TreeLevel = 0
	query := "INSERT INTO t1.tree(title) VALUES($1) RETURNING id"
	row := r.db.QueryRow(query, m.Title)

	var id int
	if err := row.Scan(&id); err != nil {
		log.WarnWithCode(3001, err.Error())
		return 0, err
	}
	return id, nil
}

func (r TreePostgres) GetAll() (*[]model.Tree, error) {
	var arr []model.Tree
	query := "SELECT * FROM t1.tree"
	err := r.db.Select(&arr, query)
	if err != nil {
		return nil, err
	}
	return &arr, nil
}

func (r TreePostgres) GetById(id int) (*model.Tree, error) {
	var task model.Tree
	query := "SELECT * FROM t1.tree WHERE id=$1"

	err := r.db.Get(&task, query, id)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return &task, err
	default:
		return &task, nil
	}
}

func (r TreePostgres) DeleteById(id int) error {
	query := "DELETE FROM t1.tree WHERE id=$1"
	_, err := r.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func NewTreePostgres(db *sqlx.DB) *TreePostgres {
	return &TreePostgres{db: db}
}
