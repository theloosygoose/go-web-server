package handler

import (
	"github.com/a-h/templ"
    "net/http"
)

func render(w http.ResponseWriter, r *http.Request , component templ.Component) error {

    return component.Render(r.Context(), w)
}
