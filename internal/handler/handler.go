package handler

import (
	"log"

	"github.com/theloosygoose/goserver/tools"
)

func CreateHandlers(queries *tools.Queries) (AdminHandler, PhotoHandler, CollectionHandler) {
    aHandler := AdminHandler{
        queries,
    }
    cHandler := CollectionHandler{
        queries,
    }
    pHandler := PhotoHandler{
        queries,
    }
    log.Println("Created Handlers")
    return aHandler, pHandler, cHandler
}
