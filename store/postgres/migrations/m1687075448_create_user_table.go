package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func m1687075448CreateUserTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1687075448",
		Migrate: func(tx *gorm.DB) error {
			type user struct {
				gorm.Model
				ID           uuid.UUID `gorm:"type:uuid;primaryKey;"`
				Email        string    `gorm:"type:varchar(255);unique"`
				Name         string    `gorm:"type:varchar(255)"`
				Username     string    `gorm:"type:varchar(255);unique"`
				Provider     string
				SocialGoogle string
				SocialFB     string
				SocialIG     string
				SocialX      string
			}
			return tx.Migrator().CreateTable(&user{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	}
}
