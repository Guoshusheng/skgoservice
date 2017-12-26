package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../datastore"
	"github.com/goji/context"
	"github.com/zenazn/goji/web"
)

func GetLoginGuo(c web.C, w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "HELLO MAN, Guo Shusheng!")
	var ctx = context.FromC(c)
	var id int64 = 1
	testsk, _ := datastore.GetTestSKID(ctx, id)
	fmt.Fprint(w, testsk)
}

func GetLoginXin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HELLO MAN, Xin Tingkai!")
	json.NewEncoder(w).Encode(
		struct {
			Name string `json:"name"`
		}{
			Name: "guo shusheng",
		})
}
