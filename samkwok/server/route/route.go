package route

import (
	"github.com/zenazn/goji/web"

	"../login"
)

func New() *web.Mux {
	mux := web.New()
	mux.Get("/gss/guoshusheng", handler.GetLoginGuo)
	mux.Post("/gss/guoshusheng", handler.GetLoginGuo)

	mux.Get("/gss/xintingkai", handler.GetLoginXin)

	return mux
}
