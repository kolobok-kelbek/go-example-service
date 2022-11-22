package entity

import "gorm.io/gorm"

type Login struct {
	gorm.Model
	Login    string
	Password string
}
