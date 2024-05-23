package handler

import (
	"net/http"

	"log"

	"github.com/theloosygoose/goserver/internal/view/components"
	"github.com/theloosygoose/goserver/tools"
)

type CollectionHandler struct {
    Queries *tools.Queries
}

func (h CollectionHandler) CreateCollection() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        new_name := r.FormValue("collection-name")

        err := h.Queries.CreateCollection(r.Context(), new_name)
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

        res, err := h.Queries.GetAllCollections(r.Context())
        if err != nil {
            log.Println("Error Getting Collections from database")
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        render(w, r, components.CollectionTable(res))
    })
}

func (h CollectionHandler) SingleCollection() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        res, err := h.Queries.GetCollectionPhotos(r.Context())
        if err != nil {
            log.Println("Error Getting Collections from database")
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        //Add A Render
        log.Println(res)

    })
}
