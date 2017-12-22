package handler

import (
	"fmt"
	"net/http"
	"time"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Enter SetCookie()")

	expiration := time.Now()
	expiration = expiration.Add(time.Duration(1) * time.Hour)
	cookie := http.Cookie{Name: "susen_go", Value: "123456", Expires: expiration}
	http.SetCookie(w, &cookie)

}

func ReadCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Enter ReadCookie()")

	cookie, _ := r.Cookie("name")
	fmt.Fprint(w, "cookie_name->"+cookie.Name)

}
