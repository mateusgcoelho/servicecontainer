package shared

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var (
	Db *sql.DB
)

func InitSqliteDatabase() {
	dbOpenned, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		panic(err)
	}
	if err = dbOpenned.Ping(); err != nil {
		panic(err)
	}
	Db = dbOpenned
	fmt.Println("Banco de dados Sqlite iniciado com sucesso.")

	runMigrations()
}

func runMigrations() {
	query := `
		CREATE TABLE IF NOT EXISTS services(
		  id INT PRIMARY KEY,
		  tag TEXT,
		  prefixUrl TEXT,
		  defaultPort INT,
		  displayName TEXT,
		  fileName TEXT,
		  engineType TEXT
		);
	`

	_, err := Db.Exec(query)

	if err != nil {
		panic(err)
	}
}
