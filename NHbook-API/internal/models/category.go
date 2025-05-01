package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"not null;unique"`
	Status      int    `gorm:"default:0"`
	Description string `gorm:"default:''"`
	Books       []Book `gorm:"foreignKey:CategoryID"`
}
