package main

import (
	"io"
	"log"
	"net"
)

const (
	host = ""
	port = "8080"
)

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
		io.WriteString(conn, `I see you connected`)
		conn.Close()
	}
}
