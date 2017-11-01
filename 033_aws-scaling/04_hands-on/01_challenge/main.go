package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {

}

func main() {
	db, err = sql.Open("mysql", "mytest:mytestpw@tcp(mytest.clkanylb1nsr.eu-central-1.rds.amazonaws.com:3306)/testdb?charset=utf8")
	check(err)
	defer db.Close()

	check(db.Ping())

	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instance)
	http.HandleFunc("/checkdb", checkdb)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":80", nil)
	check(err)
}

func checkdb(w http.ResponseWriter, req *http.Request) {
	inst := getInstance()
	io.WriteString(w, string(inst))

	rows, err := db.Query(`SELECT * from amigos;`)
	check(err)
	defer rows.Close()

	var s, name string
	s = "RETRIEVED RECORDS:\n"

	for rows.Next() {
		check(rows.Scan(&name))
		s += name + "\n"

	}
	// just for curiosity to compare both ways to write:
	io.WriteString(w, s)
	fmt.Fprintln(w, s)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from AWS.")
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK")
}

func instance(w http.ResponseWriter, req *http.Request) {
	bs := getInstance()
	io.WriteString(w, string(bs))
}

func getInstance() []byte {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	// check(err)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)
	resp.Body.Close()
	return bs
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}
