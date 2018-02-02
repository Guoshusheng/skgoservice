// a http server written in golang
// by sam kwok
// on 2017-10-26
// at changning shanghai
package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"./server/datastore"
	"./server/datastore/database"
	"./server/route"

	"code.google.com/p/go.net/context"

	"github.com/GeertJohan/go.rice"
	"github.com/drone/config"
	"github.com/drone/drone/server/middleware"
	webcontext "github.com/goji/context"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zenazn/goji/web"

	//	log "github.com/Sirupsen/logrus"
)

var (
	port = config.String("server-port", ":8089")

	// database driver configuration.
	datasource = config.String("database-datasource", "./sk.sqlite")
	driver     = config.String("database-driver", "sqlite3")

	sslcrt = config.String("server-ssl-cert", "")
	sslkey = config.String("server-ssl-key", "")

	assets_folder = config.String("server-assets-folder", "")

	db *sql.DB
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("I'm golang!")
	fmt.Println("time.Now()->", time.Now())

	expiration := time.Now()
	expiration = expiration.AddDate(0, 1, 0)
	duration := expiration.Add(time.Duration(70) * time.Second)
	fmt.Println("增加日期：expiration->", expiration)
	fmt.Println("增加时间：duration->", duration)

	// -----------------------------------------------------------------------------
	//	// 数据库设置
	//	db, err := sql.Open("sqlite3", "./drone.sqlite")
	//	checkErr(err)

	//	// 查询数据
	//	vali_code_msgs, err := db.Query("SELECT * FROM vali_code_msgs")
	//	checkErr(err)
	//	log.Infoln("vali_code_msgs", vali_code_msgs)

	//	msg, _ := vali_code_msgs.Columns()
	//	for i := range msg {
	//		fmt.Print(msg[i])
	//		fmt.Print("\t")
	//	}

	//	//删除数据
	//	stmt, err := db.Prepare("delete from customers where customer_id=?")
	//	checkErr(err)

	//	res, err := stmt.Exec(52)
	//	checkErr(err)

	//	affect, err := res.RowsAffected()
	//	checkErr(err)

	//	fmt.Println("--affect--", affect)

	//	db.Close()
	// -----------------------------------------------------------------------------

	// 数据库连接
	db = database.MustConnect(*driver, *datasource)

	// 静态文件路径设置
	var assetserve http.Handler
	if *assets_folder != "" {
		assetserve = http.FileServer(http.Dir(*assets_folder))
	} else {
		assetserve = http.FileServer(rice.MustFindBox("app").HTTPBox())
	}

	http.Handle("/robots.txt", assetserve)
	http.Handle("/static/", http.StripPrefix("/static", assetserve))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/"
		assetserve.ServeHTTP(w, r)
	})

	// 路由设置
	mux := route.New()
	mux.Use(middleware.Options)
	mux.Use(ContextMiddleware)
	mux.Use(middleware.SetHeaders)
	mux.Use(middleware.SetUser)
	http.Handle("/gss/", mux)

	// 判断http还是https访问
	if len(*sslcrt) == 0 {
		panic(http.ListenAndServe(*port, nil))
	} else {
		panic(http.ListenAndServeTLS(*port, *sslcrt, *sslkey, nil))
	}

}

// ContextMiddleware creates a new go.net/context and
// injects into the current goji context.
func ContextMiddleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		ctx = datastore.NewContext(ctx, database.NewDatastore(db))

		// add the context to the goji web context
		webcontext.Set(c, ctx)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// 统一错误处理
func checkErr(err error) {
	if err != nil {
		fmt.Println("--err---", err)
	}
}
