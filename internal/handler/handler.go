package handler

import (
	"log"

	"github.com/theloosygoose/goserver/tools"
)

func CreateHandlers(queries *tools.Queries) (AdminHandler, PhotoHandler, CollectionHandler, FormHandler) {
    a := AdminHandler{
        queries,
    }
    c := CollectionHandler{
        queries,
    }
    p := PhotoHandler{
        queries,
    }
    f := FormHandler{
        queries,
    }
    log.Println("Created Handlers")

    return a,p,c,f
}
