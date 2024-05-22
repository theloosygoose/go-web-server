package handler

import (
    "net/http"
    "database/sql"

	"log"

	"github.com/theloosygoose/goserver/internal/types"
	"github.com/theloosygoose/goserver/internal/view/components"
)

type CollectionHandler struct {
	DB *sql.DB
}

func (h CollectionHandler) CreateCollection() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        query := `INSERT INTO collections (name) VALUES(?);`

        _, err := h.DB.Exec(query, r.FormValue("collection-name"))
        if err != nil {
            log.Println("Error Creating Collection", err)
            w.WriteHeader(http.StatusInternalServerError)
        }
        w.WriteHeader(http.StatusCreated)

    })
}

func (h CollectionHandler) NewCollectionForm() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        render(w, r, components.CollectionForm())
    })
}

func (h CollectionHandler) DeleteCollection() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    })
}

func (h CollectionHandler) ShowCollectionsTable() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        query := `SELECT * FROM collections;`
        res, err := h.DB.Query(query)
        if err != nil {
            log.Println("Error Getting Collections from database")
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        var collections []types.Collection
		for res.Next() {
			var c types.Collection

			err = res.Scan(&c.ID, &c.Name, &c.Photos)
			if err != nil {
				log.Println("Failed to Scan", err)
			}

			collections = append(collections, c)
		}

        render(w, r, components.CollectionTable(collections))
    })
}

func (h CollectionHandler) SingleCollection() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        query := `SELECT img. FROM collections AS collec INNER JOIN;`
        res, err := h.DB.Query(query)
        if err != nil {
            log.Println("Error Getting Collections from database")
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        var collections []types.Collection
		for res.Next() {
			var c types.Collection

			err = res.Scan(&c.ID, &c.Name, &c.Photos)
			if err != nil {
				log.Println("Failed to Scan", err)
			}

			collections = append(collections, c)
		}

        render(w, r, components.CollectionTable(collections))
    })
}
