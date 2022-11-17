package web

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Address  string
	WhatsApp string
}

type Person struct {
	Name    string
	Age     uint
	Hobbies []string
	Info    Info
}

func Route21(w http.ResponseWriter, r *http.Request) {
	person := Person{
		Name:    "Nawi",
		Age:     22,
		Hobbies: []string{"baca novel", "main game", "baca manga"},
		Info: Info{
			Address:  "Amuntai",
			WhatsApp: "08080808",
		},
	}

	t, er := template.ParseFiles("./temp/kedua/test.html")
	if er != nil {
		http.Error(w, er.Error(), 500)
		fmt.Println(er.Error())
		return
	}
	t.Execute(w, person)
}

// agar bisa mengakses nilai yang ada di metod mana method tidak boleh pointer
func (s Person) Perkenalan() string {
	return "Hello my name is " + s.Name
}

func (s Person) Say(name string) string {
	return "Hello " + name
}

func (s Info) Perkenalan() string {
	return "I am from " + s.Address
}

func (s Info) Say(name string) string {
	return "Where are you from ," + name + "?"
}
