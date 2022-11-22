package entity

import (
	"gorm.io/gorm"
	"time"
)

type Tomato struct {
	gorm.Model
	UserId   uint
	TypeTime string
	From     time.Time
	To       time.Time
}
