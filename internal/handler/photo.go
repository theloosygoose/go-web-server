package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/theloosygoose/goserver/internal/view/photo"
	"github.com/theloosygoose/goserver/internal/view/components"
	"github.com/theloosygoose/goserver/tools"
)

type PhotoHandler struct {
    Queries *tools.Queries
}

func (h PhotoHandler) MainPage() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        res_p, err := h.Queries.GetAllPhotos(r.Context())
        if err != nil {
            log.Println("Error Running GetAllPhotos Query: ", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        res_c, err := h.Queries.GetAllCollections(r.Context())
        if err != nil{
            log.Println("Error Running GetAllCollections Query: ", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

		render(w, r, photo.Index(res_p, res_c))
	})
}

func (h PhotoHandler) ShowMainPhoto() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        id, err := strconv.Atoi(r.PathValue("id"))
        if err != nil {
            log.Println("Could not parse Path Params as Int", err)
            http.Redirect(w, r, "/", http.StatusPermanentRedirect)
            return
        }
        res, err := h.Queries.GetPhotoById(r.Context(), int64(id))

        render(w, r, photo.MainPhoto(res))
    })
}

func (h PhotoHandler) RandomPhotoShow() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        res, err := h.Queries.GetRandomPhoto(r.Context())
        if err != nil {
            log.Println("Could not Get Random Photo from DB", err)
        }

        singlePhoto := tools.Photo{
            ID: res.ID,
            Name: res.Name,
            Location: res.Location,
            Date: res.Date,
            Description: res.Description,
            Imagepath: res.Imagepath,
            IHeight: res.IHeight,
            IWidth: res.IWidth,
        }

        render(w, r, photo.MainPhoto(singlePhoto))
    })
}

func (h PhotoHandler) AllPhotos() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        res_p, err := h.Queries.GetAllPhotos(r.Context())
        if err != nil {
            log.Println("Error Running GetAllPhotos Query: ", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        render(w, r, components.GalleryItems(res_p, "All Photos"))
    })
}
