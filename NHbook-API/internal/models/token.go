package models

type Token struct {
	Token  string `gorm:"size:256;unique;not null"`
	UserID string `gorm:"size:36;not null; unique"`
}
