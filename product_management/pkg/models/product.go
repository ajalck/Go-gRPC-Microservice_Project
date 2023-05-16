package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	ProductName string  `json:"product_name" gorm:"not null"`
	Stock       int32   `json:"stock" gorm:"not null"`
	Price       float32 `json:"price" gorm:"not null"`
}
