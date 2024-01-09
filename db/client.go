package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

var db *sql.DB

func init() {
	config := mysql.NewConfig()
	config.Net = "tcp"
	config.Addr = "120.24.82.106:3456"
	config.User = "littlebox"
	config.Passwd = "cTbdJbij.a#Zwow)?207"
	config.DBName = "littlebox"
	_db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		zap.S().Fatal("db init error", "err", err)
	}
	db = _db
}
