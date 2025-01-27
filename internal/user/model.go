package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email           string    `json:"email" gorm:"uniqueIndex"`
	Name            string    `json:"name"`
	Password        string    `json:"password"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
}
