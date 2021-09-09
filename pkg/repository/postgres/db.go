package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type psqlDB struct {
	readerX *sqlx.DB
	writerX *sqlx.DB
}

func NewDBFromSql(db string) *psqlDB {
	sdb, err := sql.Open("postgres", db) // open sql db
	if err != nil {
		panic(err)
	}
	return &psqlDB{
		readerX: sqlx.NewDb(sdb, "postgres"),
		writerX: sqlx.NewDb(sdb, "postgres"),
	}
}
