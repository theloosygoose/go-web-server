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
		query := `SELECT name, location, date, imagepath FROM photos;`

		results, err := h.DB.Query(query)
		if err != nil {
			log.Println("Failed to Exectue Query: ", err)
		}

		var photos []types.Photo

		for results.Next() {
			var photo types.Photo

			err = results.Scan(&photo.Name, &photo.Location, &photo.Date, &photo.ImagePath)

			if err != nil {
				log.Println("Failed to Scan", err)
			}

			photos = append(photos, photo)
		}

		render(w, r, photo.Show(photos))
	})
}
