package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/theloosygoose/goserver/internal/types"
	"github.com/theloosygoose/goserver/internal/view/admin"
	"github.com/theloosygoose/goserver/internal/view/components"
	"github.com/theloosygoose/goserver/tools"
)

type AdminHandler struct {
    Queries *tools.Queries
}

func (h AdminHandler) CreatePhoto() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(20)

		var response types.Response

		p := tools.CreatePhotoParams{
			Name:        r.FormValue("name"),
			Location:    r.FormValue("location"),
            Description: sql.NullString{String: r.FormValue("description"), Valid: true},
		}

		file, fileHeader, err := r.FormFile("imageFile")
		if err != nil {
			log.Println(err)
			return
		}
		render(w, r, components.ReponseShow(response))
		imageProcess(file, fileHeader, &p)
		log.Println("---FILE UPLOAD COMPLETE---")

        results, err := h.Queries.CreatePhoto(r.Context(), p)
        if err != nil {
            log.Println("Could not add New Photo to Database: ", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        } else {
            log.Println("Added Row at: ", results)
        }

	})

}

func (h AdminHandler) HandlerAdminShow() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        results, err := h.Queries.GetAllCollections(r.Context())
        if err != nil {
            log.Println("Unable to get Collections")
        }

		render(w, r, admin.Show(results))
	})
}

func (h AdminHandler) DeletePhoto() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id_string := r.PathValue("id")
        id, err := strconv.Atoi(id_string)

		p, err := h.Queries.DeletePhoto(r.Context(), int64(id))
		if err != nil {
            log.Println("Unable to Delete Photo From DB: ", err)
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

		results, err := h.Queries.GetAllPhotos(r.Context())
		if err != nil {
			log.Println("Failed to Exectue Query: ", err)
		}

		render(w, r, admin.Delete(results))
	})
}
