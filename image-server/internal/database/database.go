package database

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	gormMySQL "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var db *Database
var tz *time.Location

func New() *Database {
	if db != nil {
		return db
	}

	connect()
	return db
}

func connect() {
	var err error
	if tz, err = time.LoadLocation("Asia/Tokyo"); err != nil {
		panic(err)
	}

	config := &mysql.Config{
		User:                 "testuser",
		Passwd:               "testuser",
		Net:                  "tcp",
		Addr:                 "mysql:3306",
		DBName:               "image_server",
		AllowNativePasswords: true,
		ParseTime:            true,
		Loc:                  tz,
	}

	var sqlDB *sql.DB
	if sqlDB, err = sql.Open("mysql", config.FormatDSN()); err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	var gormDB *gorm.DB
	if gormDB, err = gorm.Open(gormMySQL.New(gormMySQL.Config{
		Conn: sqlDB,
	}), &gorm.Config{}); err != nil {
		panic(err)
	}

	db = &Database{gormDB}
}
