package infrastructure

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func InitializeDatabase() (*sqlx.DB, error) {
	log.Println("[INFO] Initializing Database")

	driver := GetConfig("database.driver")
	databaseFile := GetConfig("database.file")

	connection, err := sql.Open(driver, databaseFile)
	if err != nil {
		return nil, err
	}

	database := sqlx.NewDb(connection, driver)

	log.Println("[INFO] Database Initialized Successfully")

	return database, nil
}
