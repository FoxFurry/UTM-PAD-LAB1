package database

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitDB(cfg mysql.Config) (*sqlx.DB, error) {
	return sqlx.Connect("mysql", cfg.FormatDSN())
}
