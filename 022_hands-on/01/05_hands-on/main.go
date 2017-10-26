package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

type person struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request) {
	text := `<h1>Hello index</h1>`
	w.Write([]byte(text)) // just trying to use w directly
}
func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<h1>Woof!</h1>`)
}
func me(w http.ResponseWriter, r *http.Request) {
	data := person{
		Name: `Chris`,
		Age:  42,
	}
	err := tpl.Execute(w, data)
	if err != nil {
		log.Fatalln("error executing template", err)
	}

	// io.WriteString(w, `<p>Chris</p>`)
}
func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)
}
