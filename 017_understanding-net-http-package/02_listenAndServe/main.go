package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var d hotdog
	// check for interface implementation
	var _ http.Handler = (*hotdog)(nil)
	http.ListenAndServe(":8080", d)
}
