package main

import (
	"io"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	text := `<h1>Hello index</h1>`
	w.Write([]byte(text)) // just trying to use w directly
}
func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<h1>Woof!</h1>`)
}
func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<p>Chris</p>`)
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
