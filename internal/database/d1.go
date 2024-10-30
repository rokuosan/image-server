package database

import (
	"database/sql"

	"github.com/rokuosan/image-server/internal/logger"
	_ "github.com/syumai/workers/cloudflare/d1"
)

var db *sql.DB

func init() {
	var err error
	log := logger.New()
	if db == nil {
		log.Debugf("Opening database connection")
		db, err = sql.Open("d1", "D1_DATABASE")
		if err != nil {
			log.Error(err.Error())
			panic(err)
		}
		log.Debugf("Database connection opened")
	}
}

func New() *sql.DB {
	return db
}
