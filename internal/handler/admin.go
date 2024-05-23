package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/theloosygoose/goserver/internal/types"
	"github.com/theloosygoose/goserver/internal/view/admin"
	"github.com/theloosygoose/goserver/internal/view/components"
	"github.com/theloosygoose/goserver/tools"
)

type AdminHandler struct {
	Ctx context.Context 
    Queries *tools.Queries
}

func (h AdminHandler) CreatePhoto() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(20)

		var response types.Response

		details := types.Photo{
			Name:        r.FormValue("name"),
			Location:    r.FormValue("location"),
			Description: r.FormValue("description"),
		}

		file, fileHeader, err := r.FormFile("imageFile")
		if err != nil {
			log.Println(err)
			return
		}

		render(w, r, components.ReponseShow(response))
		imageProcess(file, fileHeader, &details)
		log.Println("---FILE UPLOAD COMPLETE---")

		query := `INSERT INTO photos 
        (name, location, date, description, imagepath, i_height, i_width)
        VALUES(?, ?, ?, ?, ?, ?, ?);`

		results, err := h.DB.Exec(query,
			&details.Name, &details.Location, &details.Date, &details.Description,
			&details.Image.FileName, &details.Image.Height, &details.Image.Width)
		if err != nil {
			log.Println("Failed to Exectue Query: ", err)
			response.Message = "Failed to Execute Query"
			response.Code = http.StatusInternalServerError
		} else {
			response.Message = "Successful"
			response.Code = http.StatusOK
		}

		log.Println(results.RowsAffected())

	})

}

func (h AdminHandler) HandlerAdminShow() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        query := `SELECT * FROM collections;`
        var collections []types.Collection

        results, err := h.DB.Query(query)
        if err != nil {
            log.Println("Unable to get tags")
        }

		for results.Next() {
			var collection types.Collection

			err = results.Scan(&collection.ID, &collection.Name)
			if err != nil {
				log.Println("Failed to Scan", err)
			}

			collections = append(collections, collection)
		}

		render(w, r, admin.Show(collections))
	})
}

func (h AdminHandler) DeletePhoto() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		query := `DELETE FROM photos WHERE id = ? RETURNING imagepath;`

		var p string

		results := h.DB.QueryRow(query, id)
		err := results.Scan(&p)
		if err != nil {
			log.Println("Unable to Delete Photo", err)
		}

		cmd := exec.Command("sudo", "rm", "-rf", fmt.Sprintf("*%v", p))
		cmd.Dir = "/mnt/usb/images"

		err = cmd.Run()
		if err != nil {
			log.Println("UNABLE TO DELETE ROW")
			log.Println(err)
		}

	})
}

func (h AdminHandler) PhotoRemoveGalleryShow() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := `SELECT id, date, imagepath, i_height, i_width FROM photos;`

		results, err := h.DB.Query(query)
		if err != nil {
			log.Println("Failed to Exectue Query: ", err)
		}

		var photos []types.Photo

		for results.Next() {
			var photo types.Photo

			err = results.Scan(&photo.ID, &photo.Date, &photo.Image.FileName, &photo.Image.Height, &photo.Image.Width)

			if err != nil {
				log.Println("Failed to Scan", err)
			}

			photos = append(photos, photo)
		}

		render(w, r, admin.Delete(photos))
	})
}
