package handler

import (
	"log"
	"net/http"
	"strconv"

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
        id_string := r.PathValue("id")

        id, err := strconv.Atoi(id_string)
        if err != nil {
            log.Println("Param is not Valid")
            http.Redirect(w, r, "/admin", http.StatusPermanentRedirect)
            return
        }

        err = h.Queries.ClearCategoryCollections(r.Context(), int64(id))
        if err != nil {
            log.Println("Could not Clear Collection Photos", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        err = h.Queries.DeleteCategory(r.Context(), int64(id))
        if err != nil {
            log.Println("Could not Delete from Collection", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusOK)
    })
}

func (h CategoryHandler) ViewCategories() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        id_string := r.PathValue("id")
        id, err := strconv.Atoi(id_string)
        if err != nil{
            log.Println("Could not turn Path Value into int", err)
            http.Redirect(w,r, "/", http.StatusPermanentRedirect)
            return
        }

        res, err := h.Queries.GetCollectionPhotos(r.Context(), int64(id))
        if err != nil {
            log.Println("Error Getting Collections from database")
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        pr := collectionRowstoPhoto(res)

        render(w, r, components.GalleryItems(pr, res[0].Name_2))
    })

}
