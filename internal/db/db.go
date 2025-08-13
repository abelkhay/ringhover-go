package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Open(dsn string) (*sqlx.DB, error) {
	if dsn == "" {
		return nil, nil
	}
	return sqlx.Connect("mysql", dsn)
}
