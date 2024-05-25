package handler

import (
	"net/http"
	"strconv"

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

        res, err := h.Queries.CreateCollection(r.Context(), new_name)
        if err != nil {
            log.Println("Error Creating Collection", err)
            w.WriteHeader(http.StatusInternalServerError)
        }
        w.WriteHeader(http.StatusCreated)

        render(w, r, components.CollectionTableItem(res))
    })
}

func (h CollectionHandler) NewCollectionForm() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        render(w, r, components.CollectionForm())
    })
}

func (h CollectionHandler) DeleteCollection() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        id_string := r.PathValue("id")
        id, err := strconv.Atoi(id_string)
        if err != nil {
            log.Println("Param is not Valid")
            http.Redirect(w, r, "/admin", http.StatusPermanentRedirect)
            return
        }

        err = h.Queries.DeleteCollection(r.Context(), int64(id))

        if err != nil {
            log.Println("Could not Delete from Collection", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusOK)
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
        log.Println("Loaded photos: ", len(res))

        pr := collectionRowstoPhoto(res)

        render(w, r, components.GalleryItems(pr))
    })
}
