package main

import (
	"flag"
	"log"
	"os"

	"github.com/pressly/goose"
	"github.com/spf13/viper"

	migrationrunner "github.com/MichaelCapo23/basebuilder/migration/migrationrunner"
	_ "github.com/MichaelCapo23/basebuilder/migration/migrations"
	_ "github.com/lib/pq"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "./migrations", "directory with migration files")
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 1 {
		flags.Usage()
		return
	}
	command := args[0]

	viper.SetDefault("PG_WRITER_URI", "postgres://root:postgres@localhost/localauth?sslmode=disable")
	viper.AutomaticEnv()
	pgURL := viper.GetString("PG_WRITER_URI")
	db, err := goose.OpenDBWithDriver("postgres", pgURL)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}
	migrationrunner.RunMigration(command, db, "", arguments...)
}
