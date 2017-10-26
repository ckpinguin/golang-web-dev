package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

const (
	host = ""
	port = "8080"
)

func communicate(conn net.Conn) {
	defer conn.Close()
	// io.WriteString(conn, `I see you connected`)

	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		ln := sc.Text()
		log.Println(ln)
		// io.WriteString(conn, "you said: "+ln+"\n")
		if ln == "" { // suffices HTTP end of header or body
			break
		}
	}
	response(conn)
}

func response(conn net.Conn) {
	// just experimenting with buffer string
	body := `<html><head></head><body><h1>Hello there!</h1></body></html>`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	// fmt.Fprintf(buf)
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
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
		go communicate(conn)
	}
}
