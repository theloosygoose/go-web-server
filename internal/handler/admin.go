package handler

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/theloosygoose/goserver/internal/types"
	"github.com/theloosygoose/goserver/internal/view/admin"
    "github.com/containers/podman/v5/pkg/ctime"
)

type AdminHandler struct {
	DB *sql.DB
}

func (h AdminHandler) AdminAddPhoto() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        r.ParseMultipartForm(20)

		details := types.Photo{
			Name:      r.FormValue("name"),
			Location:  r.FormValue("location"),
		}

        file, fileHeader, err := r.FormFile("imageFile")
        if err != nil {
            log.Println(err)
            return
        }
        defer file.Close()
        log.Println("---UPLOADING PHOTO---")
        log.Printf("Name:: %v\n", fileHeader.Filename)
        log.Printf("Size:: %v\n", fileHeader.Size)

        contentType := fileHeader.Header["Content-Type"][0]
        log.Printf("Content-Type:: %v\n", contentType)

        var osFile *os.File

        if contentType == "image/jpeg"{
            osFile, err = os.CreateTemp("/mnt/usb/images/", "*.jpg")
        } else {
            osFile, err = os.CreateTemp("/mnt/usb/etc/", "")
        }
        log.Println("Error In Content-Type", err)
        defer osFile.Close()

        // SAVE FILE
        fileBytes, err := io.ReadAll(file)
        if err != nil {
            fmt.Println(err)
        }
        osFile.Write(fileBytes)
        defer osFile.Close()

        details.ImagePath = filepath.Base(osFile.Name())

        s, err := osFile.Stat()
        if err != nil{
            fmt.Println(err)
        }
        details.Date = ctime.Created(s).String()
        

        log.Println("---FILE UPLOAD COMPLETE---")

		query := `INSERT INTO photos 
        (name, location, date, imagepath, avaliable)
        VALUES($1, $2, $3, $4, $5);`

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
