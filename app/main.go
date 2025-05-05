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

	var response string

	if urlPath == "/" {
		response = "HTTP/1.1 200 OK\r\n\r\n"
	} else if strings.HasPrefix(urlPath, "/echo/") {
		response = echo(urlPath)
	} else {
		response = "HTTP/1.1 404 Not Found\r\n\r\n"
	}

	// Responding with a basic HTTP response

	conn.Write([]byte(response))
}

func echo(body string) string {
	ans := strings.TrimPrefix(body, "/echo/")
	return fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(ans), ans)
}
