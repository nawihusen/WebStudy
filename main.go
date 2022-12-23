package main

import (
	"html/template"
	"net/http"
	"web/web"
)

func main() {
	http.HandleFunc("/route21", web.Route21)
	http.HandleFunc("/route22", web.Route22)
	http.HandleFunc("/route31", web.Route31)
	http.HandleFunc("/route32", web.Route32)
	// http.HandleFunc("/route33", web.Route33) belum bisa
	http.HandleFunc("/test", web.Route31)
	http.HandleFunc("/coba", web.Route32)
	http.HandleFunc("/twelveone", web.TwelveOne)
	http.HandleFunc("/twelvetwo", web.TwelveTwo)

	var funcMap = template.FuncMap{
		"unescape": func(s string) template.HTML {
			return template.HTML(s)
		},
		"avg": func(n ...int) int {
			var total = 0
			for _, each := range n {
				total += each
			}
			return total / len(n)
		},
	}

	http.HandleFunc("/route23", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.New("view.html").
			Funcs(funcMap).
			ParseFiles("./temp/ketiga/view.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic("Gagal pada peluncuran")
	}
}
