package models

type Permission struct {
	ID             string `gorm:"column:permission_id;size:36;primaryKey"`
	PermissionName string `gorm:"size:20;unique;not null"`
}

func (p *Permission) TableName() string {
	return "permissions"
}
