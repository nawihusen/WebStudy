package web

import (
	"html/template"
	"net/http"
)

func TwelveOne(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp := template.Must(template.New("form").ParseFiles("./temp/twelve/twelve1.html"))
		err := temp.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func TwelveTwo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		temp := template.Must(template.New("result").ParseFiles("./temp/twelve/twelve2.html"))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		name := r.FormValue("name")
		msg := r.FormValue("message")

		data := map[string]string{"name": name, "message": msg}

		if er := temp.Execute(w, data); er != nil {
			http.Error(w, er.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
