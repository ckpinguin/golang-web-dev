package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getOrSetCookie(w, req)

	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer mf.Close()
		// SHA1 for filename
		ext := strings.Split(fh.Filename, ".")[1]
		hash := sha1.New()
		io.Copy(hash, mf)
		fname := fmt.Sprintf("%x", hash.Sum(nil)) + "." + ext
		log.Println("Hashed filename:", fname)

		// store on server
		pwd, err := os.Getwd()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		dst, err := os.Create(filepath.Join(pwd, "public", "pics", fname))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		mf.Seek(0, 0) // important!
		_, err = io.Copy(dst, mf)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("Saved file", fname, "on server.")
		c = appendValue(w, c, fname)
	}

	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs[1:])
}

func getOrSetCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err == http.ErrNoCookie {
		log.Println("No cookie set, we construct one...")
		sid := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sid.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	} else {
		log.Println("Picture", fname, "already is on server.")
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}
