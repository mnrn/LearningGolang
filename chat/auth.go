package main

import (
	"net/http"
	"strings"
	"fmt"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		// 未承認だった場合
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		log.Info("authHandler.serveHTTP: 未認証です。")
	} else if err != nil {
		// 別の何らかのエラーが発生
		panic(err.Error())
	} else {
		// 認証に成功した場合、ラップされたハンドラを呼び出す。
		h.next.ServeHTTP(w, r)
		log.Info("認証に成功しました。")
	}
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

// サードパーティへのログイン処理を受け付けます。
// パスの形式: /auth/{action}/{provider}
func loginHandler(w http.ResponseWriter, r *http.Request) {

	log.Info("ログインをハンドルしました。")

	segs := strings.Split(r.URL.Path, "/")
	if len(segs) < 4 {
		log.Fatal("不正なパスを読み込みました。: ", r.URL.Path)
		return
	}
	action 		:= segs[2]
	provider	:= segs[3]

	switch action {
	case "login":
		log.Info("TODO: ログイン処理", provider)
	default:
		w.WriteHeader(http.StatusNotFound)  // 404を返します。
		fmt.Fprintf(w, "アクション%sには未対応です。", action)
		log.Warningf("アクション%sには未対応です。", action)
	}
}