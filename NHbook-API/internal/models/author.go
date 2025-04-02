package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Bio       string
	BirthDate time.Time
	Books     []Book `gorm:"many2many:book_author;"`
}
