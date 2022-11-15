package main

import (
	"fmt"
	"html/template"
	"net/http"
	"web/web"
)

func main() {
	http.HandleFunc("/", web.Route1) // merupakan direct yang di tuju ketika halaman yang di
	//minta tidak terdaftar
	http.HandleFunc("/route1", web.Route1)
	http.HandleFunc("/route2", web.Route2)
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("asset"))))
	// membungkus path ke /image/
	// semua file yang ada di dalam dir asset akan di tampilkan. pemisah antar file atau folder
	// akan secara otomatis menggunakan / tergantung sistem operasi
	// http.Handle("/dir", http.FileServer(http.Dir("asset")))
	http.HandleFunc("/test", web.Route3)

	tmp, ers := template.ParseGlob("temp/*")
	// fungsi parseglob akan memparse semua file yang match dengan kondisi yang diberikan
	fmt.Println(tmp)
	if ers != nil {
		panic(ers.Error())
	}

	http.HandleFunc("/parsial1", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"title": "Belajar",
			"for":   "Golang",
			"for2":  "Sukse",
		}

		er := tmp.ExecuteTemplate(w, "index.html", data)
		if er != nil {
			http.Error(w, er.Error(), 500)
		}
	})

	http.HandleFunc("/parsial2", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"tambah": "ini tambahan",
		}
		// tinggal execute saja karena sudah di parse
		er := tmp.ExecuteTemplate(w, "baru.html", data)
		if er != nil {
			http.Error(w, er.Error(), 500)
		}
	})

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic("Gagal pada peluncuran")
	}
}

/*
kita juga bisa menggunakan yang dibawah untuk melakukan listen

var address = ":9000"
fmt.Printf("server started at %s\n", address)

server := new(http.Server)
server.Addr = address
err := server.ListenAndServe()
if err != nil {
    fmt.Println(err.Error())
}
*/
