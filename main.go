package main

import (
	_ "embed"
	"io"
	"net/http"
	"qndxx/api"
)

//go:embed index.html
var indexHtml string

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		_, _ = io.WriteString(w, indexHtml)
	}
	return
}
func main() {
	http.HandleFunc("/", handler.Handler)
	http.HandleFunc("/api/q", handler.Handler)
	_ = http.ListenAndServe(":8090", nil)
}
