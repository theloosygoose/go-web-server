package handler

import (
	"log"
	"net/http"

	"github.com/theloosygoose/goserver/internal/view/components"
	"github.com/theloosygoose/goserver/tools"
)

type CategoryHandler struct {
    Queries *tools.Queries
}

func (h CategoryHandler) NewCategory() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        new_name := r.FormValue("collection-name")

        res, err := h.Queries.CreateCategory(r.Context(), new_name)
        if err != nil {
            log.Println("Error Creating Collection", err)
            w.WriteHeader(http.StatusInternalServerError)
        }
        w.WriteHeader(http.StatusCreated)

        render(w, r, components.CollectionTableItem(res))
    })
} 

func (h CategoryHandler) DeleteCategory() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {


    })
}

func (h CategoryHandler) EditCategory() http.Handler{
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    })
}

func (h CategoryHandler) ViewCategories() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    })
}
