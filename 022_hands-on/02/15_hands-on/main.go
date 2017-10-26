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
	for sc.Scan() {
		ln := sc.Text()
		if i == 0 {
			method := strings.Fields(ln)[0]
			uri := strings.Fields(ln)[1]
			log.Println("method:", method, "uri:", uri)
		}
		if ln == "" { // suffices HTTP end of header or body
			break
		}
		i++
	}
	response(conn)
}

func response(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Code Crime</title>
		</head>
		<body>
			<h1>Hello there!</h1>
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
