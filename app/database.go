package app

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"

	"github.com/alfianyulianto/go-room-managament/config"
	"github.com/alfianyulianto/go-room-managament/halpers"
)

func NewDB() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true",
		config.Cfg.DatabaseUser,
		config.Cfg.DatabasePass,
		config.Cfg.DatabaseHost,
		config.Cfg.DatabasePort,
		config.Cfg.DatabaseName,
	)
	db, err := sql.Open("mysql", dsn)
	halpers.IfPanicError(err)

	//set database pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db
}
