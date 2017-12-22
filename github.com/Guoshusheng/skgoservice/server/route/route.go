package route

import (
	"github.com/zenazn/goji/web"

	"../handle"
)

func New() *web.Mux {
	mux := web.New()
	mux.Get("/gss/guoshusheng", handler.GetLoginGuo)
	mux.Post("/gss/guoshusheng", handler.GetLoginGuo)

	// 上传文件
	mux.Get("/gss/uploadfile", handler.UploadFile)
	mux.Post("/gss/uploadfile", handler.UploadFile)

	// 读存cookie、session
	mux.Get("/gss/login", handler.ReadCookie)
	mux.Post("/gss/login", handler.ReadCookie)

	// 测试连接
	mux.Get("/gss/xintingkai", handler.GetLoginXin)

	return mux
}
