package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}
func main() {
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
