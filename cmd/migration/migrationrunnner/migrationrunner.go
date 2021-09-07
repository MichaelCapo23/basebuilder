package migrationrunner

import (
	"database/sql"
	"log"
	"path/filepath"
	"runtime"

	goose "github.com/pressly/goose"

	_ "github.com/MichaelCapo23/jwtserver/cmd/migration/migrations"
	_ "github.com/lib/pq"
)

func RunMigration(command string, db *sql.DB, dir string, args ...string) {
	if dir == "" {
		_, b, _, _ := runtime.Caller(0)
		dir = filepath.Join(filepath.Dir(filepath.Dir(b)), "migrations")
	}

	if err := goose.Run(command, db, dir, args...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
