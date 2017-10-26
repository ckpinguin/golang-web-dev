package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const (
	host = ""
	port = "8080"
)

func handle(conn net.Conn) {
	defer conn.Close()
	sc := bufio.NewScanner(conn)
	i := 0
	var method, uri string
	inBody := false
	// headerEnd := false
	var rBody string
	for sc.Scan() {
		ln := sc.Text()
		if i == 0 {
			method = strings.Fields(ln)[0]
			uri = strings.Fields(ln)[1]
			log.Println("method:", method, "uri:", uri)
		}
		if ln == "" && method == "GET" {
			log.Println("END of GET header reached")
			break
		}
		if ln == "" && method == "POST" { // suffices HTTP end of header or body
			log.Println("END of POST header reached")
			// headerEnd = true
			inBody = true
			continue
			// rBody += ln
			// break
		}
		if method == "POST" && inBody {
			if ln == "" {
				log.Println("END of POST body reached", rBody)
				break
			}
			rBody += ln
		}

		i++
	}
	mux(conn, method, uri)
}

func mux(conn net.Conn, method string, uri string) {
	switch uri {
	case "/":
		index(conn)
	case "/apply":
		if method == "GET" {
			apply(conn)
		} else {
			procApply(conn)
		}
	default:
		index(conn)
	}
}

// let duplication arise first ;-)
func index(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Code Crime</title>
		</head>
		<body>
			<h1>Hello Home!</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
		</body>
	</html>`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}
func apply(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Code Crime</title>
		</head>
		<body>
			<h1>Apply here!</h1>
			<form action="/apply" method="POST" >
				<input type="text" name="firstName" id="fName" placeholder="Your first name">
				<input type="text" name="lastName" id="lName" placeholder="Your last name">
				<button type="submit">Save!</button>
			</form>
		</body>
	</html>`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}
func procApply(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Code Crime</title>
		</head>
		<body>
			<h1>Got an application from:</h1>
			<h2>What do you want to do now?</h2>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
		</body>
	</html>`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func main() {
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handle(conn)
	}
}
