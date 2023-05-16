package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID      int32   `json:"user_id" gorm:"not null"`
	ProductID   int32   `json:"product_id" gorm:"not null"`
	Quantity    int32   `json:"quantity" gorm:"not null"`
	TotalPrice  float32 `json:"total_price" gorm:"not null"`
	OrderStatus string  `json:"order_status"`
}
