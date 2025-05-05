package main

import (
	"fmt"
	"strings"
)

type Handler interface {
	// validate that this handle handles this url
	Accept(url string) bool
	// return t
	HandleRequest(url string) HttpResponse
}

type RootHandler struct{}

func (r RootHandler) Accept(url string) bool {
	return url == "/"
}

func (r RootHandler) HandleRequest(url string) HttpResponse {
	return RootResponse
}

type NotFoundHandler struct{}

func (n NotFoundHandler) Accept(url string) bool {
	return true
}

func (n NotFoundHandler) HandleRequest(url string) HttpResponse {
	return NotFoundResponse
}

func reformatResponse(httpResponse HttpResponse) string {
	var builder strings.Builder

	builder.WriteString("HTTP/1.1")
	if httpResponse.StatusCode == 404 {
		builder.WriteString(" 404 Not Found")
	} else if httpResponse.StatusCode == 200 {
		builder.WriteString(" 200 OK")
	}
	builder.WriteString("\r\n")
	if httpResponse.ContentType != "" {
		builder.WriteString(fmt.Sprintf("Content-Type: %s", httpResponse.ContentType))
	}
	builder.WriteString("\r\n")

	if httpResponse.ContentLength > 0 {
		builder.WriteString(fmt.Sprintf("Content-Length: %d", httpResponse.ContentLength))
	}
	builder.WriteString("\r\n")
	builder.WriteString("\r\n")

	builder.WriteString(httpResponse.Body)

	return builder.String()
}
