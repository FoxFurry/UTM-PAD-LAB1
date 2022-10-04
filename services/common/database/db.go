package database

import (
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func InitDB(ctx context.Context, cfg mysql.Config) (*sql.DB, error) {
	return sql.Open("mysql", cfg.FormatDSN())
}
