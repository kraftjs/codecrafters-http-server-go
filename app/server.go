package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	okayResponse     = "HTTP/1.1 200 OK\r\n\r\n"
	notFoundResponse = "HTTP/1.1 404 Not Found\r\n\r\n"
)

func handleConnection(conn net.Conn) {
	request := readRequest(conn)
	sendResponse(conn, request)
	conn.Close()
}

func readRequest(conn net.Conn) Request {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection: ", err.Error())
	}
	return NewRequest(buf)
}

func sendResponse(conn net.Conn, request Request) {
	var response string
	switch {
	case request.path == "/":
		response = okayResponse
	case strings.HasPrefix(request.path, "/echo/"):
		response = CreateEchoResponse(request.path[len("/echo/"):])
	default:
		response = notFoundResponse
	}
	fmt.Println("Response: ", response)
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
}
