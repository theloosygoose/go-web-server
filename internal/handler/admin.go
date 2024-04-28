package handler

import (
	"database/sql"
	"github.com/theloosygoose/goserver/internal/types"
	"github.com/theloosygoose/goserver/internal/view/admin"
	"log"
	"net/http"
)

type AdminHandler struct {
	DB *sql.DB
}

func (h AdminHandler) AdminAddPhoto() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := `INSERT INTO photos 
        (name, location, date, imagepath, avaliable)
        VALUES($1, $2, $3, $4, $5);`

		details := types.Photo{
			Name:      r.FormValue("name"),
			Location:  r.FormValue("location"),
			Date:      r.FormValue("date"),
			ImagePath: r.FormValue("imagepath"),
		}

		results, err := h.DB.Exec(query, &details.Name, &details.Location, &details.Date, &details.ImagePath, true)
		if err != nil {
			log.Println("Failed to Exectue Query: ", err)
		}

		log.Println(results.RowsAffected())
	})

}

func (h AdminHandler) HandlerAdminShow() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		render(w, r, admin.Show())
	})
}
