package register

import (
	"database/sql"

	"github.com/vnlozan/goose/v3"
)

func init() {
	goose.AddMigrationNoTx(
		func(_ *sql.DB) error { return nil },
		func(_ *sql.DB) error { return nil },
	)
}
