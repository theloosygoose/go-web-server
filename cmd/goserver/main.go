package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"database/sql"

	"github.com/joho/godotenv"

	"github.com/theloosygoose/goserver/internal/routes"
	"github.com/theloosygoose/goserver/tools"
)

func main() {
    ctx := context.Background()

    err := godotenv.Load(".env")
    if err != nil {
        log.Println("Loaded Env File")
    }
    db, err := sql.Open("sqlite3", os.Getenv("DB_CONNINFO"))
    queries := tools.New(db)

	r := routes.NewServer(ctx, queries)

	nerr := http.ListenAndServe(":8080", r)
	if nerr != nil {
		log.Fatal(nerr)
	}

}
