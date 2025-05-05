package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	//Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", ":4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	handlers := []Handler{
		RootHandler{},
		EchoHandler{},
		NotFoundHandler{},
	}

	reader := bufio.NewReader(conn)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}

	// The first line of an HTTP request looks like: "GET /abcdefg HTTP/1.1"
	parts := strings.Fields(requestLine)
	if len(parts) < 2 {
		fmt.Println("Invalid request format")
		return
	}

	urlPath := parts[1] // Extract the requested path
	fmt.Println("User requested:", urlPath)

	for _, handler := range handlers {
		if handler.Accept(urlPath) {
			conn.Write([]byte(reformatResponse(handler.HandleRequest(urlPath))))
			break
		}
	}
}

type EchoHandler struct{}

func (e EchoHandler) Accept(url string) bool {
	return strings.HasPrefix(url, "/echo/")
}

func (e EchoHandler) HandleRequest(url string) HttpResponse {
	return CreateHttpResponse(200, "text/plain", strings.TrimPrefix(url, "/echo/"))
}
