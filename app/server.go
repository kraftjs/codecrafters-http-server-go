package main

import (
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	const response = "HTTP/1.1 200 OK\r\n\r\n"
	conn.Write([]byte(response))
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	connection, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	handleConnection(connection)
	connection.Close()
}
