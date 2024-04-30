package routes

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/theloosygoose/goserver/internal/handler"
)

func NewServer(db *sql.DB) *http.ServeMux {
	r := http.NewServeMux()

    pHandler := handler.PhotoHandler{DB: db}
    aHandler := handler.AdminHandler{DB: db}
    dist := os.Getenv("STATIC_DIR")

    fs := http.FileServer(http.Dir(dist))

    r.Handle("/dist/", http.StripPrefix("/dist/", fs))

	r.HandleFunc("/", pHandler.HandlerPhotoShow())
    r.HandleFunc("/addphoto", aHandler.AdminAddPhoto())

    r.HandleFunc("/admin", aHandler.HandlerAdminShow())


	return r
}
