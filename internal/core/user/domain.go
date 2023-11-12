package user

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Username     string    `json:"username"`
	Provider     string    `json:"provider"`
	SocialGoogle string    `json:"socialGoogle"`
	SocialFB     string    `json:"socialFb"`
	SocialIG     string    `json:"socialIg"`
	SocialX      string    `json:"socialX"`
}
