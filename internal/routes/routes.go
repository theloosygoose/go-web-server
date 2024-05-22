package routes

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/theloosygoose/goserver/internal/handler"
)

func NewServer(db *sql.DB) *http.ServeMux {
	r := http.NewServeMux()

    pHandler := handler.PhotoHandler{DB: db}
    aHandler := handler.AdminHandler{DB: db}
    cHandler := handler.CollectionHandler{DB: db}

    dist := os.Getenv("STATIC_DIR")
    log.Println(dist)
    fs := http.FileServer(http.Dir(dist))
    r.Handle("/dist/", http.StripPrefix("/dist/", fs))

	r.HandleFunc("/", pHandler.HandlerPhotoShow())
    r.HandleFunc("GET /photodata/{id}", pHandler.HandlerMainPhotoShow())
    r.HandleFunc("GET /photodata/random", pHandler.HandlerRandomPhotoShow())

    r.HandleFunc("PUT /addphoto", aHandler.CreatePhoto())
    r.HandleFunc("GET /admin", aHandler.HandlerAdminShow())

    r.HandleFunc("GET /photodata/delete", aHandler.PhotoRemoveGalleryShow())
    r.HandleFunc("DELETE /photodata/{id}", aHandler.DeletePhoto())


    r.HandleFunc("GET /collections", cHandler.ShowCollectionsTable())
    r.HandleFunc("POST /collections", cHandler.CreateCollection())
    r.HandleFunc("GET /collections/form", cHandler.NewCollectionForm())

    r.HandleFunc("GET /collections/{id}", cHandler.SingleCollection())
    r.HandleFunc("DELETE /collections/{id}", cHandler.DeleteCollection())

	return r
}
