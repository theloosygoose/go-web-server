package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/theloosygoose/goserver/internal/types"
	"github.com/theloosygoose/goserver/internal/view/components"
	"github.com/theloosygoose/goserver/tools"
)

type FormHandler struct {
    Queries *tools.Queries
}


func (h FormHandler) NewForm() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        res, err := h.Queries.GetAllCollections(r.Context())
        if err != nil {
            log.Println("Could not Get collections from DB New Form", err)
        }

        v := types.FormwithValuesEmpty(res)

        render(w, r, components.NewForm(v))
    })
}

func (h FormHandler) UpdateForm() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        id_string := r.PathValue("id")

        id, err := strconv.Atoi(id_string)
        if err != nil {
            log.Println("Could not convert pathvalue to int: ", err)
        }

        res_p, err := h.Queries.GetPhotoById(r.Context(), int64(id))
        if err != nil {
            log.Println("Could not Get Photo by ID UpdateForm: ", err)
        }

        res_c, err := h.Queries.GetAllCollections(r.Context())
        if err != nil {
            log.Println("Could Not Get Collections from DB UpdateForm: ", err)
        }

        res_pc, err := h.Queries.PhotoIDGetCollections(r.Context(), int64(id))
        if err != nil {
            log.Println("Could Not Get Photos Collections from DB PhotoIdGetCollections : ", err)
        }
        var m map[string]string

        for _, c := range res_c {
            r := c.Name + fmt.Sprint(c.ID)
            m[r] = "off"
        }

        for _, c := range res_pc {
            r := c.Name + fmt.Sprint(c.ID)
            m[r] = "on"
        }

        v := types.FormwithValues(res_c, m, res_p)

        render(w, r, components.UpdateForm(v ))
        
    })
}
