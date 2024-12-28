package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // SQLite driver
)

var DB *sql.DB

func InitializeDatabase() {
	var err error
	DB, err = sql.Open("sqlite", "./myproducts.db") // Change the DB name to myproducts.db
	if err != nil {
		log.Fatal(err)
	}

	// Create Products table if not exists
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		price REAL,
		quantity INTEGER
	);`)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return DB
}
