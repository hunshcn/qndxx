package main

import (
	"net/http"
	"qndxx/api"
)

func main() {
	http.HandleFunc("/q", handler.Handler)
	_ = http.ListenAndServe(":8090", nil)
}
