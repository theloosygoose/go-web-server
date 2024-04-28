package routes

import (
	"database/sql"
	"net/http"

	"github.com/theloosygoose/goserver/internal/handler"
)

func NewServer(db *sql.DB) *http.ServeMux {
	r := http.NewServeMux()

    pHandler := handler.PhotoHandler{DB: db}
    aHandler := handler.AdminHandler{DB: db}

	r.HandleFunc("GET /", pHandler.HandlerPhotoShow())
    r.HandleFunc("POST /addphoto", aHandler.AdminAddPhoto())

    r.HandleFunc("GET /admin", aHandler.HandlerAdminShow())

	return r
}
