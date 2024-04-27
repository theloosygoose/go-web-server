package db

import (
	"database/sql"
	"fmt"
	"log"
    _ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "dadphoto"
)

type PGdb struct {*sql.DB}

func NewDB(sqldb *sql.DB) PGdb {
    return PGdb{sqldb}
}

func Connect() *sql.DB {
	connInfo := fmt.Sprintf("host=%s port=%d \n user=%s password=%s \n dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Panicln(err)
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
		results, err := db.Query("CREATE TABLE photos (id SERIAL PRIMARY KEY, name VARCHAR(100) NOT NULL, location VARCHAR(100) NOT NULL, date VARCHAR(100), imagepath VARCHAR(100) NOT NULL, avaliable BOOL);")
		if err != nil {
			fmt.Println("failed to execute query", err)
			return
		}
		fmt.Println("Table created successfully", results)

	} else {

		fmt.Println("Table 'photos' already exists ")
	}
}
