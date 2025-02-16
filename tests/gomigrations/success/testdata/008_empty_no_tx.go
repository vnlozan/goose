package gomigrations

import (
	"github.com/vnlozan/goose/v3"
)

func init() {
	goose.AddMigrationNoTx(nil, nil)
}
