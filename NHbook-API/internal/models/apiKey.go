package models

import "gorm.io/gorm"

type ApiKey struct {
	gorm.Model
	ApiKey string `gorm:"unique;not null"`
	Status uint   `gorm:"default:1"`
}
