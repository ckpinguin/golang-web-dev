package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"path"
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

	route := request(conn)
	fmt.Println("route: ", path.Clean(route))
	mux(conn, path.Clean(route))
}

func mux(conn net.Conn, route string) {
	var content string
	if route == "/index.html" {
		content = `<h1>hello there</h1>`
		respond(conn, content)
	} else {
		respond404(conn)
	}

}

func request(conn net.Conn) string {
	i := 0 // line number tracking
	scanner := bufio.NewScanner(conn)
	var uri string
	for scanner.Scan() {
		ln := scanner.Text()
		// fmt.Println(ln)
		if i == 0 {
			method := strings.Fields(ln)[0]
			uri = strings.Fields(ln)[1]
			fmt.Println("*METHOD*:", method, "*URI*:", uri)
		}
		if ln == "" {
			break
		}
		i++
	}
	return uri
}

func respond404(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Document</title>
	</head>
	<body>
		<h1>404 Not found that site!</h1>
	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 404 NOTFOUND\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	// Same here, empty line means "end of header stuff, body follows"
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func respond(conn net.Conn, content string) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Document</title>
	</head>
	<body>
		` + content +
		`</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	// Same here, empty line means "end of header stuff, body follows"
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
