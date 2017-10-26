// a http server written in golang
// by sam kwok
// on 2017-10-26
// at changning shanghai
package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"./samkwok/server/route"

	"github.com/GeertJohan/go.rice"
	"github.com/drone/config"

	log "github.com/Sirupsen/logrus"
	_ "github.com/mattn/go-sqlite3"
)

var (
	port = config.String("server-port", ":8089")

	sslcrt = config.String("server-ssl-cert", "")
	sslkey = config.String("server-ssl-key", "")

	assets_folder = config.String("server-assets-folder", "")
)

func main() {
	fmt.Println("Hello World!")

	// 数据库设置
	db, err := sql.Open("sqlite3", "./drone.sqlite")
	checkErr(err)

	// 查询数据
	vali_code_msgs, err := db.Query("SELECT * FROM vali_code_msgs")
	checkErr(err)
	log.Infoln("vali_code_msgs", vali_code_msgs)

	msg, _ := vali_code_msgs.Columns()
	for i := range msg {
		fmt.Print(msg[i])
		fmt.Print("\t")
	}

	//	for vali_code_msgs.Next() {
	//		var msg_id int
	//		var phone string
	//		var code string
	//		var typecode string
	//		var send_time string
	//		var send_count int

	//		err = vali_code_msgs.Scan(&msg_id, &phone, &code, &typecode, &send_time, &send_count)
	//		checkErr(err)
	//		fmt.Println(phone)
	//		fmt.Println(send_time)
	//	}

	//删除数据
	stmt, err := db.Prepare("delete from customers where customer_id=?")
	checkErr(err)

	res, err := stmt.Exec(52)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("--affect--", affect)

	db.Close()

	// 静态文件路径设置
	var assetserve http.Handler
	if *assets_folder != "" {
		assetserve = http.FileServer(http.Dir(*assets_folder))
		log.Infoln("---assets_forder!=kong")
	} else {
		assetserve = http.FileServer(rice.MustFindBox("app").HTTPBox())
		log.Infoln("---assets_forder---k--")
	}

	http.Handle("/robots.txt", assetserve)
	http.Handle("/static/", http.StripPrefix("/static", assetserve))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/"
		assetserve.ServeHTTP(w, r)
	})

	log.Infoln("---The Middle---")
	fmt.Println("Come on!")

	// 路由设置
	mux := route.New()
	http.Handle("/gss/", mux)

	// 判断http还是https访问
	if len(*sslcrt) == 0 {
		panic(http.ListenAndServe(*port, nil))
	} else {
		panic(http.ListenAndServeTLS(*port, *sslcrt, *sslkey, nil))
	}

}

// 统一错误处理
func checkErr(err error) {
	if err != nil {
		fmt.Println("--err---", err)
	}
}
