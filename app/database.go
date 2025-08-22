package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"

	"github.com/alfianyulianto/go-room-managament/halpers"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_room_management?multiStatements=true&parseTime=true")
	halpers.IfPanicError(err)

	//set database pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db
}
