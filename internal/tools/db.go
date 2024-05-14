package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

    _ "github.com/mattn/go-sqlite3"
)

type PGdb struct {*sql.DB}

func NewDB(sqldb *sql.DB) PGdb {
    return PGdb{sqldb}
}

func Connect() *sql.DB {
    connInfo := os.Getenv("DB_CONNINFO")

	db, err := sql.Open("sqlite3", connInfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Connected to db!")
	return db
}

func CloseConnection(db *sql.DB) {
	defer db.Close()
}

func CreateTable(db *sql.DB) {
	var exists bool

	if err := db.QueryRow("SELECT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'photos' );").Scan(&exists); err != nil {
		fmt.Println("failed to execute query", err)
		return
	}

	if !exists {
		results, err := db.Query("CREATE TABLE photos (id SERIAL PRIMARY KEY, name VARCHAR(100) NOT NULL, location VARCHAR(100) NOT NULL, date VARCHAR(100), imagepath VARCHAR(100) NOT NULL, description TEXT, i_height VARCHAR(10), i_width VARCHAR(10));")
		if err != nil {
			fmt.Println("failed to execute query", err)
			return
		}
		fmt.Println("Table created successfully", results)

	} else {

		fmt.Println("Table 'photos' already exists ")
	}
}
