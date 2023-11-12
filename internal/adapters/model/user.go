package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
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

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		uuid := uuid.New()
		u.ID = uuid
	}
	return nil
}
