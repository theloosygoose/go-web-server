package handler

import (
	"github.com/a-h/templ"
    "net/http"
    "log"
    "os"
    "os/exec"
    "io"
    "fmt"

	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/theloosygoose/goserver/internal/types"
	"github.com/containers/podman/v5/pkg/ctime"
)

func render(w http.ResponseWriter, r *http.Request , component templ.Component) error {

    return component.Render(r.Context(), w)
}

func imageProcess(file multipart.File, header *multipart.FileHeader, i *types.Photo) error {
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
		osFile.Name(), "-resize", "150000@\\>",
		filepath.Dir(osFile.Name())+"/sm_"+i.Image.FileName)

	medcmd := exec.Command("sudo", "magick",
		osFile.Name(), "-resize", "1000000@\\>",
		filepath.Dir(osFile.Name())+"/med_"+i.Image.FileName)

	lgcmd := exec.Command("sudo", "magick",
		osFile.Name(), "-resize", "2000000@\\>",
		filepath.Dir(osFile.Name())+"/lg_"+i.Image.FileName)

	cmds = append(cmds, mincmd, smcmd, medcmd, lgcmd)
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
