package main

import (
	"log"
	"net/http"
	"os"

	"database/sql"

	"github.com/joho/godotenv"
	"github.com/theloosygoose/goserver/internal/routes"
	"github.com/theloosygoose/goserver/tools"

    _"github.com/mattn/go-sqlite3"
)

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Println("Loaded Env File")
    }

    db, err := sql.Open("sqlite3", os.Getenv("DB_CONNINFO"))
    if err != nil {
        log.Println("Error Connecting to DB", err)
    }

    queries := tools.New(db)
	r := routes.NewServer(queries)

	nerr := http.ListenAndServe(":8080", r)
	if nerr != nil {
		log.Fatal(nerr)
	}

}
