package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:text(150);not null"`
	Password string `gorm:"size:255;not null"`
}
