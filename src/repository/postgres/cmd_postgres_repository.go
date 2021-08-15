package postgres

import (
	"github.com/Askalag/piece16/src/log"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"path/filepath"
)

const (
	filePath = "src/migration/sql/"
	initFile = "initial_tables.sql"
	dropFile = "drop.sql"
)

type CmdPostgres struct {
	db *sqlx.DB
}

// InitTables create tables and schemas if not exists
func (r *CmdPostgres) InitTables() error {
	return r.executeSqlFIle(filePath, initFile)
}

// DropAll drop tables and schemas if exists
func (r *CmdPostgres) DropAll() error {
	return r.executeSqlFIle(filePath, dropFile)
}

func (r *CmdPostgres) executeSqlFIle(p string, fileName string) error {
	path := filepath.Join(p, fileName)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.WarnWithCode(1000, err.Error())
		return err
	}

	sql := string(f)
	_, err = r.db.Exec(sql)
	if err != nil {
		log.WarnWithCode(1000, err.Error())
		return err
	}
	return nil
}

func NewCmdPostgres(db *sqlx.DB) *CmdPostgres {
	return &CmdPostgres{db: db}
}
