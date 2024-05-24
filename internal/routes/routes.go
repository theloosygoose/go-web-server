package routes

import (
	"net/http"
	"os"

	"github.com/theloosygoose/goserver/internal/handler"
	"github.com/theloosygoose/goserver/tools"
)

func NewServer(
    queries *tools.Queries,

) *http.ServeMux {
	r := http.NewServeMux()

    aHandler,
    pHandler,
    cHandler := handler.CreateHandlers(queries)

    dist := os.Getenv("STATIC_DIR")
    fs := http.FileServer(http.Dir(dist))
    r.Handle("/dist/", http.StripPrefix("/dist/", fs))

	r.HandleFunc("/", pHandler.MainPage())

    r.HandleFunc("GET /photodata", pHandler.AllPhotos())
    r.HandleFunc("GET /photodata/{id}", pHandler.ShowMainPhoto())
    r.HandleFunc("GET /photodata/random", pHandler.RandomPhotoShow())

    r.HandleFunc("PUT /addphoto", aHandler.CreatePhoto())
    r.HandleFunc("GET /admin", aHandler.AdminShow())

    r.HandleFunc("GET /photodata/delete", aHandler.PhotoRemoveGalleryShow())
    r.HandleFunc("DELETE /photodata/{id}", aHandler.DeletePhoto())


    r.HandleFunc("GET /collections", cHandler.ShowCollectionsTable())
    r.HandleFunc("PUT /collections", cHandler.CreateCollection())

    r.HandleFunc("GET /collections/{id}", cHandler.SingleCollection())
    r.HandleFunc("DELETE /collections/{id}", cHandler.DeleteCollection())

	return r
}
