package main

import (
	"net/http"
	"web/web"
)

func main() {
	http.HandleFunc("/route21", web.Route21)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic("Gagal pada peluncuran")
	}
}
