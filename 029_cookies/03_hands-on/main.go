package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("visits")
	if err == http.ErrNoCookie { // no cookie, first visit?
		cookie = &http.Cookie{
			Name:  "visits",
			Value: "0",
		}
	}

	visits, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	visits++
	cookie.Value = strconv.Itoa(visits)

	http.SetCookie(w, cookie)

	fmt.Fprintln(w, "You were here", visits, "times.")
}
