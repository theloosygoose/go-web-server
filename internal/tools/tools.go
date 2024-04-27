package tools

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

func ConnectDB(){
    connStr := "postgres://postgres:password@localhost:5432/db-dad?sslmode=disable"

    db, err := sql.Open("postgres", connStr)
    defer db.Close()
    if err != nil{
        log.Fatal(err)
    }

    if err = db.Ping(); err != nil{
        log.Fatal(err)
    }
}
