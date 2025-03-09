// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
package initializer

import (
	"fmt"
	"time"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() {
	configDB := global.Config.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", configDB.User, configDB.Password, configDB.Host, configDB.Port, configDB.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	if err != nil {
		panic(fmt.Sprintf("Open MYSQL ERROR %v", err))
	}

	global.MySQL = db
	setPool()
	global.Logger.Info("Connect MySQL Success", zap.String("msg", "Success"))
}

func setPool() {
	configMYSQL := global.Config.MySQL
	db, err := global.MySQL.DB()
	if err != nil {
		panic(fmt.Sprintf("Setup MySQL Error: %v", err))
	}
	db.SetMaxOpenConns(configMYSQL.MaxOpenConnect)
	db.SetMaxIdleConns(configMYSQL.MaxIdleConnect)
	db.SetConnMaxLifetime(time.Duration(configMYSQL.MaxConnectTimeLife))
}
