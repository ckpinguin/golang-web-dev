package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", foo)
	http.Handle("/dog/", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func bar(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bar ran")
}