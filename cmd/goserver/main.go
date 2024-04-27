package main

import (
	"log"
	"net/http"

	"github.com/theloosygoose/goserver/internal/routes"
	db "github.com/theloosygoose/goserver/internal/tools"
)

func main() {
    DB := db.Connect()
    db.CreateTable(DB)

	r := routes.NewServer(DB)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

    db.CloseConnection(DB)
}
