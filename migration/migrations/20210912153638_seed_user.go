package migrations

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upSeedUser, downSeedUser)
}

func upSeedUser(tx *sql.Tx) error {
	_, err := tx.Exec(fmt.Sprintf(`INSERT INTO "user" (id, external_id, email, first_name, last_name) VALUES ('%s', '%s', '%s', '%s', '%s');`, uuid.New(), "WSHIIlUNgYQ0ZJuZGBByrlzulTB3", "michael@airvet.com", "Michael", "Capobianco"))
	if err != nil {
		return err
	}
	return nil
}

func downSeedUser(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DELETE * FROM "user";`)
	if err != nil {
		return err
	}
	return nil
}
