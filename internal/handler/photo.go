package handler

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/theloosygoose/goserver/internal/types"
	"github.com/theloosygoose/goserver/internal/view/photo"
	"github.com/theloosygoose/goserver/tools"
)

type PhotoHandler struct {
	Ctx context.Context 
    Queries *tools.Queries
}

func (h PhotoHandler) HandlerPhotoShowAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        results, err := h.Queries.GetAllPhotos(h.Ctx)

        if err != nil {
            log.Println("Error Running GetAllPhotos Query: ", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

		render(w, r, photo.PhotoCard(results))
	})
}

func (h PhotoHandler) HandlerMainPhotoShow() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        id, err := strconv.Atoi(r.PathValue("id"))
        if err != nil {
            log.Println("Could not parse Path Params as Int", err)
            http.Redirect(w, r, "/", http.StatusPermanentRedirect)
            return
        }
        results, err := h.Queries.GetPhotoById(h.Ctx, int64(id))

        render(w, r, photo.MainPhoto(results))
    })
}

func (h PhotoHandler) HandlerRandomPhotoShow() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        res, err := h.Queries.GetRandomPhoto(h.Ctx)
        if err != nil {
            log.Println("Could not Get Random Photo from DB", err)
        }

        render(w, r, photo.MainPhoto(res))
    })
}
