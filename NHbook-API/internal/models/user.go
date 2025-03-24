package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"column:user_id;primaryKey;size:36"`
	UserName  string `gorm:"size:20;not null"`
	Phone     string `gorm:"size:10"`
	Email     string `gorm:"size:50;unique;not null"`
	Password  string `gorm:"size:100"`
	Avatar    string `gorm:"size:255"`
	Status    int    `gorm:"default:1"`
	Address   string `gorm:"size:100"`
	Roles     []Role `gorm:"many2many:user_roles"`
	Token     Token  `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) TableName() string {
	return "users"
}
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	return nil
}
