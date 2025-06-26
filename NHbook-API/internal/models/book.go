package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string    `gorm:"not null"`
	Authors     []Author  `gorm:"many2many:book_author;"`
	ImageURL    string    `gorm:"default:''"`
	Price       int       `gorm:"default:0"`
	Description string    `gorm:"type:text"`
	Stock       int       `gorm:"default:0"`
	CategoryID  int       `gorm:"not null"`
	Category    Category  `gorm:"foreignKey:CategoryID"`
	PublishedAt time.Time `gorm:"type:date"`
}
