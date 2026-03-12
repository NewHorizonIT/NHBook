package database

import (
	"database/sql"
	"fmt"

	"github.com/NewHorizonIT/nhbook-api/internal/shared/config"
	_ "github.com/lib/pq"
)

func InitDB(cnf *config.DbConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", GetDNS(cnf))
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cnf.MaxIdle)
	db.SetMaxOpenConns(cnf.MaxOpen)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func GetDNS(cnf *config.DbConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cnf.Host, cnf.Port, cnf.User, cnf.Password, cnf.Name)
}
