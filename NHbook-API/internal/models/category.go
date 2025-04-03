package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"not null;unique"`
	Description string `gorm:"default:''"`
	Books       []Book `gorm:"foreignKey:CategoryID"`
}
