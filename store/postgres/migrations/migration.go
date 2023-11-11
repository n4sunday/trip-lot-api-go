package migrations

import (
	"errors"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	m := gormigrate.New(
		db,
		gormigrate.DefaultOptions,
		[]*gormigrate.Migration{
			// Add migration here
		},
	)

	if err := m.Migrate(); err != nil {
		if errors.Is(err, gormigrate.ErrNoMigrationDefined) {
			log.Println("No migration to run")
			return
		}
		log.Fatalf("Could not migrate: %v", err)
	}
}
