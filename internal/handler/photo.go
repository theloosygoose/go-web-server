package handler

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/theloosygoose/goserver/internal/types"
	"github.com/theloosygoose/goserver/internal/view/photo"
)

type PhotoHandler struct {
	DB *sql.DB
}

func (h PhotoHandler) HandlerPhotoShow() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := `SELECT id, name, location, date, imagepath FROM photos;`

		results, err := h.DB.Query(query)
		if err != nil {
			log.Println("Failed to Exectue Query: ", err)
		}

		var photos []types.Photo

		for results.Next() {
			var photo types.Photo

			err = results.Scan(&photo.ID, &photo.Name, &photo.Location, &photo.Date, &photo.ImagePath)

			if err != nil {
				log.Println("Failed to Scan", err)
			}

			photos = append(photos, photo)
		}

		render(w, r, photo.PhotoCard(photos))
	})
}

func (h PhotoHandler) HandlerMainPhotoShow() http.HandlerFunc{
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        id := r.PathValue("id") 
        query := `SELECT (id, name, location, date, imagepath) 
        FROM photos 
        WHERE id = $1;`

        results := h.DB.QueryRow(query, id)

        var p types.Photo
        err := results.Scan(&p.ID, &p.Name, &p.Location, &p.Date, &p.ImagePath)
        if err != nil {
            log.Println("Main Photo not Found",  err)
        }

        render(w, r, photo.MainPhoto(p))
    })
}
