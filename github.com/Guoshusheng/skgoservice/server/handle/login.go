package handler

import (
	"fmt"
	"net/http"
)

func GetLoginGuo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HELLO MAN, Guo Shusheng!")
}

func GetLoginXin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HELLO MAN, Xin Tingkai!")
}
