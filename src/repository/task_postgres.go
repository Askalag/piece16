package repository

import (
	"github.com/Askalag/piece16/src/log"
	"github.com/Askalag/piece16/src/model"
	"github.com/jmoiron/sqlx"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func (repo *TaskPostgres) Create(obj model.Task) (int64, error) {
	res, err := repo.db.NamedExec(
		`INSERT INTO task (title, tree_level) VALUES (:title, :treeLevel)`,
		map[string]interface{}{
			"title": obj.Title,
			"treeLevel": obj.TreeLevel,
		})
	if err != nil {
		log.WarnWithCode(3001, err.Error())
		return 0, err
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		log.WarnWithCode(3001, err.Error())
	}
	return lastInsertId, nil
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}