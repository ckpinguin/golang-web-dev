package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	// A connections always starts with a request
	request(conn)

	// and it finishes with the server's response
	respond(conn)
}

func request(conn net.Conn) {
	i := 0 // line number tracking
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		// First line is the request itself
		if i == 0 {
			method := strings.Fields(ln)[0]
			uri := strings.Fields(ln)[1]
			fmt.Println("*METHOD*:", method, "*URI*:", uri)
		}
		// empty line is "end of header" in http/1.1 standard
		if ln == "" {
			break
		}
		i++
	}
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Document</title>
	</head>
	<body>
		<h1>Hello there :-)</h1>
	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	// Same here, empty line means "end of header stuff, body follows"
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
