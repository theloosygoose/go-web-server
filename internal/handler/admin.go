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
	"strings"

	"github.com/containers/podman/v5/pkg/ctime"
	"github.com/theloosygoose/goserver/internal/types"
	"github.com/theloosygoose/goserver/internal/view/admin"
	"github.com/theloosygoose/goserver/internal/view/components"
)

type AdminHandler struct {
	DB *sql.DB
}

func (h AdminHandler) AdminAddPhoto() http.HandlerFunc {
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
		ImageProcess(file, fileHeader, &details)
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

		render(w, r, admin.Show())
	})
}

func (h AdminHandler) HandlerAdminDeletePhoto() http.HandlerFunc {
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

func (h AdminHandler) HandlerAdminDeleteShow() http.HandlerFunc {
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
		"'%[fx:w]x%[fx:h]'",
		osFile.Name())

	size, err := sizecmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(size)
	xy := strings.TrimFunc(string(size[:]), func(r rune) bool {
		if r == '\'' || r == ' ' {
			return true
		}
		return false
	})
	xy_str := strings.Split(xy, "x")
	i.Image.Width = xy_str[0]
	i.Image.Height = xy_str[1]

	var cmds []*exec.Cmd

	fmt.Println("---RUNNING MAGICK---")
	mincmd := exec.Command("sudo", "magick",
		osFile.Name(), "-resize", "50x50",
		filepath.Dir(osFile.Name())+"/min_"+i.Image.FileName)

	smcmd := exec.Command("sudo", "magick",
		osFile.Name(), "-resize", "250000@\\>",
		filepath.Dir(osFile.Name())+"/sm_"+i.Image.FileName)

	medcmd := exec.Command("sudo", "magick",
		osFile.Name(), "-resize", "1000000@\\>",
		filepath.Dir(osFile.Name())+"/med_"+i.Image.FileName)

	cmds = append(cmds, mincmd, smcmd, medcmd)
	go magickCommand(cmds)

	defer osFile.Close()

	return nil
}

func magickCommand(cmds []*exec.Cmd) {
	for i := range cmds {
		err := cmds[i].Run()
		if err != nil {
			fmt.Println(err)
		}

	}
}
