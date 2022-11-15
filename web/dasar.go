package web

import (
	"html/template"
	"net/http"
	"path"
)

func Route1(write http.ResponseWriter, r *http.Request) {
	msg := "Ini adalah isi dari web yang kita buat"
	write.Write([]byte(msg))
}

func Route2(write http.ResponseWriter, r *http.Request) {
	msg := "Ini adalah isi dari web kedua yang kita buat"
	write.Write([]byte(msg))
}

func Route3(w http.ResponseWriter, r *http.Request) {

	filepath := path.Join("temp", "index.html")

	data := map[string]interface{}{
		"title": "Belajar Golang",
		"for":   "Meningkatkan Kemampuan belajar Golang",
		"for2":  "Menjadi Golang develover",
	}

	t, err := template.ParseFiles(filepath)
	// t, err := template.ParseFiles("./temp/view.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// jika terjadi error maka error akan masuk ke http.error

	err = t.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
