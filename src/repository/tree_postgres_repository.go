package repository

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/jmoiron/sqlx"
)

type TreePostgres struct {
	db *sqlx.DB
}

func (t TreePostgres) Create(m *model.Task) (int, error) {
	panic("implement me")
}

func (t TreePostgres) GetAll() (*[]model.Tree, error) {
	panic("implement me")
}

func (t TreePostgres) GetById(id int) (*model.Task, error) {
	panic("implement me")
}

func (t TreePostgres) DeleteById(id int) error {
	panic("implement me")
}

func NewTreePostgres(db *sqlx.DB) *TreePostgres {
	return &TreePostgres{db: db}
}
