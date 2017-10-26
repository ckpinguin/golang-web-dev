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
	fmt.Println("Code got here.")
	io.WriteString(conn, `I see you connected`)

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
