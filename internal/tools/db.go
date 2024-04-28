package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type PGdb struct {*sql.DB}

func NewDB(sqldb *sql.DB) PGdb {
    return PGdb{sqldb}
}

func Connect() *sql.DB {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error Loading .env file")
    }

    var (
        host = os.Getenv("DB_HOST")
        port = os.Getenv("DB_PORT")
        user = os.Getenv("DB_USER")
        password = os.Getenv("DB_PASS")
        dbname = os.Getenv("DB_NAME")
    )


	connInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
    log.Println(connInfo)

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
