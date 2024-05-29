package routes

import (
	"net/http"

	"github.com/theloosygoose/goserver/dist"
	"github.com/theloosygoose/goserver/internal/handler"
	"github.com/theloosygoose/goserver/tools"
)


func NewServer(
    queries *tools.Queries,

) *http.ServeMux {
	r := http.NewServeMux()

    aHandler,
    pHandler,
    cHandler,
    fHandler:= handler.CreateHandlers(queries)

    fs := http.FileServer(http.FS(dist.Files))
    r.Handle("/dist/", http.StripPrefix("/dist/", fs))

	r.HandleFunc("/", pHandler.MainPage())

    r.HandleFunc("GET /photodata", pHandler.AllPhotos())
    r.HandleFunc("GET /photodata/{id}", pHandler.ShowMainPhoto())
    r.HandleFunc("GET /photodata/random", pHandler.RandomPhotoShow())

    r.HandleFunc("PUT /photodata", aHandler.CreatePhoto())
    r.HandleFunc("GET /admin", aHandler.AdminShow())

    r.HandleFunc("GET /photodata/delete", aHandler.PhotoRemoveGalleryShow())

    r.HandleFunc("PUT /photodata/{id}", aHandler.UpdatePhoto())
    r.HandleFunc("DELETE /photodata/{id}", aHandler.DeletePhoto())

    r.HandleFunc("GET /collections", cHandler.ShowCollectionsTable())
    r.HandleFunc("PUT /collections", cHandler.CreateCollection())

    r.HandleFunc("GET /collections/{id}", cHandler.SingleCollection())
    r.HandleFunc("DELETE /collections/{id}", cHandler.DeleteCollection())

    r.HandleFunc("GET /admin/form/new", fHandler.NewForm())
    r.HandleFunc("GET /admin/form/{id}", fHandler.UpdateForm())

	return r
}
