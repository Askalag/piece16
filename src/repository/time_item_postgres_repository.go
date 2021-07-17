package repository

import "github.com/jmoiron/sqlx"

type TimeItemPostgres struct {
	db *sqlx.DB
}

func NewTimeItemPostgres(db *sqlx.DB) *TimeItemPostgres {
	return &TimeItemPostgres{db: db}
}
