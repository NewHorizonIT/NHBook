package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID  uint
	Order    Order `gorm:"foreignKey:OrderID"`
	BookID   uint
	Book     Book `gorm:"foreignKey:BookID"`
	Quantity int
	Price    int // giá lúc mua
}
