package migrations

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.Token{},
		&models.ApiKey{},
	)

	if err != nil {
		global.Logger.Error("Migrate err", zap.String("err", err.Error()))
	} else {
		global.Logger.Info("Migrate success")
	}
}
