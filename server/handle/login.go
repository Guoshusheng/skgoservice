package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLoginGuo(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "HELLO MAN, Guo Shusheng!")
	json.NewEncoder(w).Encode(
		struct {
			Name string `json:"name"`
		}{
			Name: "guo shusheng",
		})
}

func GetLoginXin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HELLO MAN, Xin Tingkai!")
}
