package main

import "net"

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, _ := ln.Accept()
		go echo(conn)
	}
}

func echo(conn net.Conn) {
	conn.Write([]byte("Hello, world"))
}

