package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	okayResponse     = ""
	notFoundResponse = ""
)

func handleConnection(conn net.Conn) {
	var buffer []byte
	conn.Read(buffer)
	response := parseRequest(string(buffer))
	conn.Write([]byte(response))
}

func parseRequest(request string) string {
	fmt.Println("Request:\n", request)
	startLine := strings.Split(request, "\r\n")[0]
	path := strings.Split(startLine, " ")[1]

	switch path {
	case "/":
		return okayResponse
	default:
		return notFoundResponse
	}
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
