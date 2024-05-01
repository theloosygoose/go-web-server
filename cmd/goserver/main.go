package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/theloosygoose/goserver/internal/routes"
	db "github.com/theloosygoose/goserver/internal/tools"
)

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Println("Loaded Env File")
    }

    DB := db.Connect()
    db.CreateTable(DB)

	r := routes.NewServer(DB)

	nerr := http.ListenAndServe(":8080", r)
	if nerr != nil {
		log.Fatal(nerr)
	}

    db.CloseConnection(DB)
}
