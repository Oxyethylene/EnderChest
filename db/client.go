package db

import (
	"database/sql"
	dbConfig "github.com/Oxyethylene/littlebox/config"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

var db *sql.DB

func init() {
	config := mysql.NewConfig()
	config.Net = dbConfig.Database.Net
	config.Addr = dbConfig.Database.Addr
	config.User = dbConfig.Database.User
	config.Passwd = dbConfig.Database.Password
	config.DBName = dbConfig.Database.Database
	dsn := config.FormatDSN()
	zap.S().Info("attempt connect to database %s", dsn)
	_db, err := sql.Open(dbConfig.Database.Driver, dsn)
	if err != nil {
		zap.S().Fatal("db init error", "err", err)
	}
	db = _db
	zap.S().Info("connect to database succeed")
}
