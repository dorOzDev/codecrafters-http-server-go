package main

import (
	"fmt"
	"strings"
)

type Handler interface {
	// validate that this handle handles this url
	accept(httpRequest HttpRequest) bool
	// return t
	handleRequest(httpRequest HttpRequest) HttpResponse
}

type RootHandler struct{}

func (r RootHandler) accept(httpRequest HttpRequest) bool {
	return httpRequest.Path() == "/"
}

func (r RootHandler) handleRequest(httpRequest HttpRequest) HttpResponse {
	return RootResponse
}

type NotFoundHandler struct{}

func (n NotFoundHandler) accept(httpRequest HttpRequest) bool {
	return true
}

func (n NotFoundHandler) handleRequest(httpRequest HttpRequest) HttpResponse {
	return NotFoundResponse
}

func reformatResponse(httpResponse HttpResponse) string {
	var builder strings.Builder

	builder.WriteString("HTTP/1.1 ")
	builder.WriteString(httpResponse.status.String())
	builder.WriteString("\r\n")
	if httpResponse.contentType != "" {
		builder.WriteString(fmt.Sprintf("Content-Type: %s", httpResponse.contentType))
	}
	builder.WriteString("\r\n")

	if httpResponse.contentLength > 0 {
		builder.WriteString(fmt.Sprintf("Content-Length: %d", httpResponse.contentLength))
	}
	builder.WriteString("\r\n")
	builder.WriteString("\r\n")

	builder.WriteString(httpResponse.body)

	return builder.String()
}
