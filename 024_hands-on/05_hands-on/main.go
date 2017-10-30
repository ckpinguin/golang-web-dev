package main

import (
	"html/template"
	"net/http"
)

var tpl template.Template

func init() {
	tpl := template.Must(template.ParseFiles("starting-files/templates/index.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	pub := http.FileServer(http.Dir("starting-files/public"))
	http.Handle("/public", pub)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
