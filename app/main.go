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

	request, _ := parseGetRequest(conn)

	handlers := []Handler{
		RootHandler{},
		EchoHandler{},
		NotFoundHandler{},
	}

	for _, handler := range handlers {
		if handler.accept(request) {
			conn.Write([]byte(reformatResponse(handler.handleRequest(request))))
			break
		}
	}
}

func parseGetRequest(conn net.Conn) (*GetRequest, error) {
	raw, err := parseRawRequest(conn)
	if err != nil {
		return nil, fmt.Errorf("failed to read request line: %w", err)
	}

	lines := strings.Split(raw, "\r\n")
	if len(lines) < 1 {
		return nil, fmt.Errorf("empty request")
	}

	parts := strings.Split(lines[0], " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid request line: %s", lines[0])
	}

	method, path, version := parts[0], parts[1], parts[2]
	if method != "GET" {
		return nil, fmt.Errorf("unsupported method: %s", method)
	}

	headers := make(map[string]string)
	for _, line := range lines[1:] {
		if line == "" {
			break
		}

		kv := strings.SplitN(line, ":", 2)
		if len(kv) == 2 {
			headers[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}

	queryParams := make(map[string]string)
	pathParts := strings.SplitN(path, "?", 2)
	cleanPath := pathParts[0]
	if len(pathParts) == 2 {
		pairs := strings.Split(pathParts[1], "&")
		for _, pair := range pairs {
			kv := strings.SplitN(pair, "=", 2)
			if len(kv) == 2 {
				queryParams[kv[0]] = kv[1]
			}
		}
	}

	return &GetRequest{
		MethodValue:  method,
		PathValue:    cleanPath,
		VersionValue: version,
		RawValue:     raw,
		Headers:      headers,
		QueryParams:  queryParams,
	}, nil
}

func parseRawRequest(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read request line: %w", err)
	}

	requestLine = strings.TrimSpace(requestLine)
	headers := ""
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("failed to read header line: %w", err)
		}

		if line == "\r\n" || line == "\n" {
			break // end of headers
		}

		headers += line
	}

	return requestLine + "\r\n" + headers + "\r\n", nil
}
