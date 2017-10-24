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
			fmt.Println(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 { // first line is the request
			mux(conn, ln)
			// A full http server would handle the rest of the
			// header here, before actually muxing
		}
		if ln == "" { // empty line means end of headers
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	method := strings.Fields(ln)[0]
	uri := strings.Fields(ln)[1]

	// multiplexer
	if method == "GET" && uri == "/" {
		index(conn)
	}
	if method == "GET" && uri == "/about" {
		about(conn)
	}
	if method == "GET" && uri == "/contact" {
		contact(conn)
	}
	if method == "GET" && uri == "/apply" {
		apply(conn)
	}
	if method == "POST" && uri == "/apply" {
		doApply(conn)
	}
}

func index(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<h1>INDEX</h1>
	<ul>
		<li><a href="/">Index</a></li>
		<li><a href="/about">About</a></li>
		<li><a href="/contact">Contact</a></li>
		<li><a href="/apply">Apply</a></li>
	</ul>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<h1>ABOUT</h1>
	<ul>
		<li><a href="/">Index</a></li>
		<li><a href="/about">About</a></li>
		<li><a href="/contact">Contact</a></li>
		<li><a href="/apply">Apply</a></li>
	</ul>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<h1>CONTACT</h1>
	<ul>
		<li><a href="/">Index</a></li>
		<li><a href="/about">About</a></li>
		<li><a href="/contact">Contact</a></li>
		<li><a href="/apply">Apply</a></li>
	</ul>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func apply(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<h1>APPLY</h1>
	<ul>
		<li><a href="/">Index</a></li>
		<li><a href="/about">About</a></li>
		<li><a href="/contact">Contact</a></li>
		<li><a href="/apply">Apply</a></li>
	</ul>
	<form action="/apply" method="post">
		<input type="text" name="Username" placeholder="Your username" id="user_name">
		<input type="password" name="Password" id="user_pw">
		<button type="submit">Submit</button>
	</form>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func doApply(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<h1>APPLICATION DONE!</h1>
	<h2>Thank you for subscribing to our spam newsletters</h2>
	<ul>
		<li><a href="/">Index</a></li>
		<li><a href="/about">About</a></li>
		<li><a href="/contact">Contact</a></li>
		<li><a href="/apply">Apply</a></li>
	</ul>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
