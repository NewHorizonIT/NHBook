package models

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	Token  string `gorm:"size:256;not null"`
	UserID string `gorm:"size:36;not null"`
}
