package auth

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	LoginId  uint
	Username string
	Email    string
}
