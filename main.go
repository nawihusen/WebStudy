package main

import (
	"net/http"
	"web/web"
)

func main() {
	http.HandleFunc("/route21", web.Route21)
	http.HandleFunc("/route22", web.Route22)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic("Gagal pada peluncuran")
	}
}
