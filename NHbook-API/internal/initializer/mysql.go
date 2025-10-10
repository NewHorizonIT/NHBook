// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
package initializer

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func InitMySQL() {
	// Step 1: Connect mysql
	configDB := global.Config.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", configDB.User, configDB.Password, configDB.Host, configDB.Port, configDB.Name)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(fmt.Sprintf("Open MYSQL ERROR %v", err))
	}
	global.MySQL = db

	// Step 2: Set Pool
	configMYSQL := global.Config.MySQL
	db.SetMaxOpenConns(configMYSQL.MaxOpenConnect)
	db.SetMaxIdleConns(configMYSQL.MaxIdleConnect)
	db.SetConnMaxLifetime(time.Duration(configMYSQL.MaxConnectTimeLife))
	global.Logger.Info("Connect MySQL Success", zap.String("msg", "Success"))

	if err := db.Ping(); err != nil {
		global.Logger.Error("Ping Database error")
	}

}
