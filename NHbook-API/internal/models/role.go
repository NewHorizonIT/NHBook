package models

type Role struct {
	ID          string       `gorm:"column:role_id;size:36;primaryKey"`
	RoleName    string       `gorm:"size:20;not null"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
}

func (r *Role) TableName() string {
	return "roles"
}
