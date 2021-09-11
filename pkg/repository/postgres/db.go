package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

type PsqlDB struct {
	ReaderX *sqlx.DB
	WriterX *sqlx.DB
}

func NewDBFromSql(db string) *PsqlDB {
	// sdb, err := sql.Open("postgres", db) // open sql db
	// if err != nil {
	// 	panic(err)
	// }
	pdb, err := goose.OpenDBWithDriver("postgres", db)
	if err != nil {
		panic(err)
	}
	return &PsqlDB{
		ReaderX: sqlx.NewDb(pdb, "postgres"),
		WriterX: sqlx.NewDb(pdb, "postgres"),
	}
}
