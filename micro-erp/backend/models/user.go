package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	name     string `gorm:"size:100;not null"`
	Email    string `gorm:"size:100;not null;unique"`
	Password string `gorm:"not null"`
	Role     string `gorm:"default:'user'"`
}
