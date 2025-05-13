package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
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
	for {
		req, err := parseHttpRequest(conn)
		if err != nil {
			fmt.Println("Closing connection due to: ", err)
			return
		}
		resp := HandleHttpRequest(req)
		conn.Write([]byte(resp.reformatResponse(req)))
		if shouldClose(req) {
			fmt.Println("Closing connection due to request by user")
			return
		}
	}
}

func shouldClose(req HttpRequest) bool {
	val, exists := req.hasHeader(CONNECTION)
	if exists {
		return strings.ToLower(val) == "close"
	}

	return false
}

func parseHttpRequest(conn net.Conn) (HttpRequest, error) {
	raw, err := readRawRequest(conn)
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
	httpMethod, isValidMethod := GetHttpMethod(method)
	if !isValidMethod {
		return nil, fmt.Errorf("unsupported method: %s", method)
	}

	headers := extractHeaders(lines)
	queryParams, cleanPath := extractQueryParams(path)

	if httpMethod.Equals(GET) {
		return &GetRequest{
			methodValue:    httpMethod,
			pathValue:      cleanPath,
			versionValue:   version,
			rawValue:       raw,
			headersMap:     headers,
			queryParamsMap: queryParams,
		}, nil
	} else if httpMethod.Equals(POST) {
		return &PostRequest{
			methodValue:    httpMethod,
			pathValue:      cleanPath,
			versionValue:   version,
			rawValue:       raw,
			headersMap:     headers,
			queryParamsMap: queryParams,
			bodyValue:      extractBody(raw),
		}, nil
	}

	return nil, fmt.Errorf("unexpected error happend, shouldn't get here")
}

func extractBody(rawRequest string) string {
	parts := strings.SplitN(rawRequest, "\r\n\r\n", 2)
	if len(parts) < 2 {
		return ""
	}
	return parts[1]
}

func readRawRequest(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)

	// Read headers until \r\n\r\n
	var rawRequest strings.Builder
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		rawRequest.WriteString(line)
		if line == "\r\n" {
			break
		}
	}

	headersPart := rawRequest.String()
	contentLength := 0
	scanner := bufio.NewScanner(strings.NewReader(headersPart))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.ToLower(line), strings.ToLower(CONTENT_LENGTH)) {
			value := strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
			cl, err := strconv.Atoi(value)
			if err != nil {
				return "", fmt.Errorf("invalid Content-Length: %v", err)
			}
			contentLength = cl
			break
		}
	}

	if contentLength > 0 {
		body := make([]byte, contentLength)
		_, err := io.ReadFull(reader, body)
		if err != nil {
			return "", fmt.Errorf("failed to read body: %v", err)
		}
		rawRequest.Write(body)
	}

	return rawRequest.String(), nil
}

func extractQueryParams(path string) (map[string]string, string) {
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
	return queryParams, cleanPath
}

func extractHeaders(lines []string) map[string]string {
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
	return headers
}
