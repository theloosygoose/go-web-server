package handler

import (
	"context"

	"github.com/theloosygoose/goserver/tools"
)

func CreateHandlers(ctx context.Context, queries *tools.Queries) (AdminHandler, PhotoHandler, CollectionHandler) {
    aHandler := AdminHandler{
        ctx,
        queries,
    }
    cHandler := CollectionHandler{
        ctx,
        queries,
    }
    pHandler := PhotoHandler{
        ctx,
        queries,
    }

    return aHandler, pHandler, cHandler
}
