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
		var res types.Response
        res.Message = "Successfully Updated Photo"
        res.Code = http.StatusOK

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


		imageProcess(file, fileHeader, &p)

		log.Println("---FILE UPLOAD COMPLETE---")

        id, err := h.Queries.CreatePhoto(r.Context(), p)
        if err != nil {
            log.Println("Could not add New Photo to Database: ", err)
            res.Message = "Could not Add photo Into Database"
            res.Code = http.StatusInternalServerError
            w.WriteHeader(http.StatusInternalServerError)
            return
        } 
        log.Println("Added Row at: ", id)

        c, err := h.Queries.GetAllCollections(r.Context())
        if err != nil {
            log.Println("Could not get all collections", err)
            res.Message = "Could not Add photo Into Database"
            res.Code = http.StatusInternalServerError
            return
        }
        log.Println("Got Collections to add new Photos")

        sc := checkBoxHandler(r, c)
        log.Println("Checked Array: ",sc)

        for _, checked := range sc {
            n := tools.PhotoIntoCollectionParams{
                PhotoID: id,
                CollectionID: checked,
            }
            log.Println("INTO COLLECTION: ", id, " " ,checked)

            err = h.Queries.PhotoIntoCollection(r.Context(), n)
            if err != nil {
                log.Println("Error Adding Photo to Collection", err)
                res.Message = "Could not Add photo Into Database"
                res.Code = http.StatusInternalServerError
            }
        }

		render(w, r, components.ReponseShow(res))
	})

}

func (h AdminHandler) AdminShow() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		render(w, r, admin.Show())
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
        ps := photoRowtoPhoto(results)

		render(w, r, admin.PhotosTable(ps))
	})
}

func (h AdminHandler) UpdatePhoto() http.HandlerFunc{
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id_string := r.PathValue("id")
        id, err := strconv.Atoi(id_string)
        var res types.Response
        res.Message = "Successfully Updated Photo"
        res.Code = http.StatusOK

		err = h.Queries.ClearPhotoCollections(r.Context(), int64(id))
		if err != nil {
            log.Println("Unable to Delete Photo From DB: ", err)
            res.Code = http.StatusInternalServerError
            res.Message = "Could not Delete Photo From DB"
		}

        path, err := h.Queries.GetPhotoPath(r.Context(), int64(id))
        if err != nil {
            log.Println("Could not get photo path: ", err)
            res.Code = http.StatusInternalServerError
            res.Message = "Could not get photo path from DB"
            return
        }


        c, err := h.Queries.GetAllCollections(r.Context())
        if err != nil {
            log.Println("Could not get all collections", err)
            res.Code = http.StatusInternalServerError
            res.Message = "Could not get all collections"
            return
        }

        sc := checkBoxHandler(r, c)

        for _, checked := range sc {
            n := tools.PhotoIntoCollectionParams{
                PhotoID: int64(id),
                CollectionID: checked,
            }
            err = h.Queries.PhotoIntoCollection(r.Context(), n)
            if err != nil {
                log.Println("Error Adding Photo to Collection", err)
                res.Code = http.StatusInternalServerError
                res.Message = fmt.Sprint("Error Adding Photo of ID: ", id, " into new collection")
            }
        }

        params := tools.UpdatePhotoParams{
			Name:        r.FormValue("name"),
			Location:    r.FormValue("location"),
            Description: sql.NullString{String: r.FormValue("description"), Valid: true},
            Imagepath: path,
            ID: int64(id),
        }

		err = h.Queries.UpdatePhoto(r.Context(), params)
		if err != nil {
            log.Println("Unable to Update From into DB: ", err)
            res.Code = http.StatusInternalServerError
            res.Message = fmt.Sprint("Unable to Update Photo into DB")
		}

        render(w, r, components.ReponseShow(res))
    })
}
