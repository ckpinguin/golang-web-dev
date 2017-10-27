package main

import (
	"log"
	"net/http"
)

func main() {
	// log.Println(http.Dir("."))
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
