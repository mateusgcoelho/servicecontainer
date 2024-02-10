package app

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitSqliteDatabase() *sql.DB {
	dbOpenned, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		panic(err)
	}
	if err = dbOpenned.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Banco de dados Sqlite iniciado com sucesso.")

	runMigrations(dbOpenned)
	return dbOpenned
}

func runMigrations(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS services(
		  id INT PRIMARY KEY,
		  tag TEXT,
		  suffixUrl TEXT,
		  defaultPort INT,
		  displayName TEXT,
		  fileName TEXT,
		  engineType TEXT
		);
	`

	_, err := db.Exec(query)

	if err != nil {
		panic(err)
	}
}
