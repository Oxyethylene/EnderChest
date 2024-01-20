package db

import (
	"database/sql"
	dbConfig "github.com/Oxyethylene/littlebox/config"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

var db *sql.DB

func InitDbClient() {
	config := mysql.NewConfig()
	config.Net = dbConfig.DatabaseConfig.Net
	config.Addr = dbConfig.DatabaseConfig.Addr
	config.User = dbConfig.DatabaseConfig.User
	config.Passwd = dbConfig.DatabaseConfig.Password
	config.DBName = dbConfig.DatabaseConfig.Database
	dsn := config.FormatDSN()
	zap.S().Infow("attempt connect to database", "dsn", dsn)
	_db, err := sql.Open(dbConfig.DatabaseConfig.Driver, dsn)
	if err != nil {
		zap.S().Fatalw("db init error", zap.Error(err))
	}
	db = _db
	zap.S().Info("connect to database succeed")
}
