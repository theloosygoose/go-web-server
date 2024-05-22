package handler

import (
    "net/http"
    "database/sql"

	"log"

	"github.com/theloosygoose/goserver/internal/types"
	"github.com/theloosygoose/goserver/internal/view/admin"
	"github.com/theloosygoose/goserver/internal/view/components"
)



type CollectionHandler struct {
	DB *sql.DB
}

func (h CollectionHandler) CreateCollection() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        query := `INSERT INTO collections (name) VALUES(?);`

        _, err := h.DB.Exec(query, r.FormValue("name"))
        if err != nil {
            log.Println("Error Creating Collection", err)
            w.WriteHeader(http.StatusInternalServerError)
        }
        w.WriteHeader(http.StatusCreated)

    })
}

func (h CollectionHandler) NewCollectionForm() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    })
}

func (h CollectionHandler) DeleteCollection() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    })
}

func (h CollectionHandler) ShowCollectionsTable() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    })
}
