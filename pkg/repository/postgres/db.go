package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type PsqlDB struct {
	ReaderX *sqlx.DB
	WriterX *sqlx.DB
}

func NewDBFromSql(db string) *PsqlDB {
	sdb, err := sql.Open("postgres", db) // open sql db
	if err != nil {
		panic(err)
	}
	return &PsqlDB{
		ReaderX: sqlx.NewDb(sdb, "postgres"),
		WriterX: sqlx.NewDb(sdb, "postgres"),
	}
}
