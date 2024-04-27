package tools

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() {
	connStr := "postgres://postgres:password@localhost:5432/dadphoto?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createPhotosTable(db)

    photo := Photo{Name: "Photo1", Location: "Home", Date: "10/21/2001", Imagepath: "/photo1_min.jpg", Avaliable: true}

    pk := insertPhoto(db, photo)

    photo1 := getFromKey(db, pk)

    fmt.Printf("ID = %d\n", pk)

    fmt.Printf("Photo Name: %s \n", photo1.Name)
    fmt.Printf("PhotoData: %v \n", photo1)
}

func createPhotosTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS photos (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    location VARCHAR(100) NOT NULL,
    date VARCHAR(100),
    imagepath VARCHAR(100) NOT NULL,
    avaliable BOOLEAN,
    created timestamp DEFAULT NOW())
    `

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func insertPhoto(db *sql.DB, photo Photo) int {
	query := `
    INSERT INTO photos (name, location, date, imagepath, avaliable)
    VALUES ($1, $2, $3, $4, $5) RETURNING id
    `

	var pk int

    err := db.QueryRow(query, photo.Name, photo.Location, photo.Date, photo.Imagepath, photo.Avaliable).Scan(&pk)
    if err != nil {
        log.Fatal(err)
    }

    return pk
}

func getFromKey(db *sql.DB, pk int) Photo {
    query :=  `SELECT name, avaliable, imagepath, date FROM photos WHERE id = $1`

    var photo Photo

    err := db.QueryRow(query, pk).Scan(&photo.Name, &photo.Avaliable, &photo.Imagepath, &photo.Date)
    if err != nil{
        log.Fatal(err)
    }

    return photo
}
