package main

import (
	"fmt"
	"slices"
	"strings"
)

type Request struct {
	httpMethod  string
	path        string
	httpVersion string
	headers     []string
	body        string
}

func NewRequest(buf []byte) Request {
	data := string(buf)
	fmt.Print("Request:\n", data)

	httpRequestLines := strings.Split(data, "\r\n")
	fmt.Print("httpRequestLines:\n", httpRequestLines)

	startLine := httpRequestLines[0]
	startLineParams := strings.Split(startLine, " ")

	var headers []string
	for i := 1; i < len(httpRequestLines); i++ {
		if httpRequestLines[i] == "" {
			break
		}
		headers = append(headers, httpRequestLines[i])
	}

	body := ""
	emptyLineIndex := slices.Index(httpRequestLines, "\r\n")
	if emptyLineIndex > 0 && emptyLineIndex < len(httpRequestLines) {
		body = strings.Join(httpRequestLines[emptyLineIndex+1:], "\r\n")
	}

	return Request{
		httpMethod:  startLineParams[0],
		path:        startLineParams[1],
		httpVersion: startLineParams[2],
		headers:     headers,
		body:        body,
	}
}

func CreateEchoResponse(bleh string) string {
	startLine := "HTTP/1.1 200 OK\r\n"
	headers := []string{
		"Content-Type: text/plain\r\n",
		fmt.Sprintf("Content-Length: %v\r\n", len(bleh)),
	}
	emptyLine := "\r\n\r\n"
	body := fmt.Sprintf("%s\r\n\r\n", bleh)

	return startLine + strings.Join(headers, "") + emptyLine + body
}
