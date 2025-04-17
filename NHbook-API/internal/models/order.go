package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID        string `gorm:"column:user_id;size:36"`
	TotalAmount   int
	Status        string
	PaymentMethod string
	OrderItems    []OrderItem
}
