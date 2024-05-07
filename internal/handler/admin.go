package handler

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/containers/podman/v5/pkg/ctime"
	"github.com/theloosygoose/goserver/internal/types"
	"github.com/theloosygoose/goserver/internal/view/admin"
)

type AdminHandler struct {
	DB *sql.DB
}

func (h AdminHandler) AdminAddPhoto() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(20)

		details := types.Photo{
			Name:     r.FormValue("name"),
			Location: r.FormValue("location"),
            Description: r.FormValue("description"),
		}

		file, fileHeader, err := r.FormFile("imageFile")
		if err != nil {
			log.Println(err)
			return
		}

        ImageProcess(file, fileHeader, &details)
		log.Println("---FILE UPLOAD COMPLETE---")

		query := `INSERT INTO photos 
        (name, location, date, imagepath)
        VALUES($1, $2, $3, $4);`

		results, err := h.DB.Exec(query, &details.Name, &details.Location, &details.Date, &details.Image.FileName)
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

func ImageProcess(file multipart.File, header *multipart.FileHeader, i *types.Photo) error {
		contentType := header.Header["Content-Type"][0]
		log.Printf("Content-Type:: %v\n", contentType)

		var osFile *os.File
        var err error

		if contentType == "image/jpeg" {
			osFile, err = os.CreateTemp("/mnt/usb/images/", "*.jpg")

		} else if contentType == "image/png" {
			osFile, err = os.CreateTemp("/mnt/usb/images/", "*.png")
		}
		if err != nil {
            return err
		}
		defer osFile.Close()

		// SAVE FILE
		fileBytes, err := io.ReadAll(file)
		if err != nil {
            return err
		}
		osFile.Write(fileBytes)
		defer osFile.Close()

        // Get File Data
		s, err := osFile.Stat()
		if err != nil {
            return err
		}

        // Get File Date
		year, month, day := ctime.Created(s).Local().Date()
		i.Date = fmt.Sprintf("%v %v, %v", year, month, day)
        

        //MAGICK EXECUTION
		i.Image.FileName = filepath.Base(osFile.Name())
        sizecmd := exec.Command("identify", "-format", 
            "'%[fx:w] x %[fx:h] pixels'",
            osFile.Name())

        size, err := sizecmd.Output()
        if err != nil{
            return err
        }
        fmt.Println(size)


        fmt.Println("---RUNNING SMALL MAGICK---")
		mincmd := exec.Command("sudo", "magick",
			osFile.Name(), "-resize", "50x50",
			filepath.Dir(osFile.Name()) + "/min_"+ i.Image.FileName)
        magickCommand(mincmd)

        defer osFile.Close()

        fmt.Println("---RUNNING MED MAGICK---")
		medcmd := exec.Command("sudo", "magick",
			osFile.Name(), "-resize", "50%",
			filepath.Dir(osFile.Name()) + "/med_"+ i.Image.FileName)

        magickCommand(medcmd)

        defer osFile.Close()

        return nil 
}

func magickCommand(cmd *exec.Cmd) {
    err := cmd.Run()
    if err != nil {
        fmt.Println(err)
    }
}
