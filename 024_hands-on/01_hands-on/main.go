package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}
func main() {
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog.jpg", dogImg)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func dogImg(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpg")
}
func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<h1>foo ran</h1>`)
}
