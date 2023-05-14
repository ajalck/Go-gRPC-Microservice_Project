package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   uint   `gorm:"primarykey,not null"`
	Email    string `gorm:"unique" binding:"required,email"`
	Password string `gorm:"unique" binding:"required,min=6"`
}
