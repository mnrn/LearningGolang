package main

import (

	"net/http"
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/gorilla/pat"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/gplus"
	"os"
	"fmt"
)


// Global variables
var (
	log = logrus.New()
	addr = flag.String("addr", ":8080", " アプリケーションのアドレス")
)

func init() {
	// フラグ解釈する。
	flag.Parse()

	// Gothのセットアップ
	goth.UseProviders(
		gplus.New(os.Getenv("GPLUS_KEY"), os.Getenv("GPLUS_SECRET"), fmt.Sprintf("http://localhost%s/auth/gplus/callback", *addr)),
	)
}


func main() {
	log.Info("main: ルーティングを開始します。")
	router := pat.New()
    router.Add("GET", "/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	router.Add("GET", "/login", &templateHandler{filename: "login.html"})
	router.Get("/auth/", loginHandler)
	r := newRoom()
	router.Add("GET","/room", r)
	log.Info("main: ルーティングを終了しました。" )

	// チャットルームを開始する。
    go r.run()

	// Webサーバーを開始する。
	log.Info("Webサーバーを開始します。ポート: ", *addr)
	log.Fatal(http.ListenAndServe(*addr, router))
}
