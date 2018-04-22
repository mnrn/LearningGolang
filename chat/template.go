package main

import (
	"sync"
	"net/http"
	"path/filepath"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template  // templは1つのテンプレートを表す。
}

// HTTPリクエストを処理する。
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	log.Info("templateHandler.ServeHTTP: HTTP接続を開始します。")
	if err := t.templ.Execute(w, r); err != nil {
		log.Error("templateHandler.ServeHTTP: ", err)
	}
}
