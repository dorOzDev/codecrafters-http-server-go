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
	return httpRequest.path() == "/"
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

func reformatResponse(httpRequest HttpRequest, httpResponse HttpResponse) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("%s ", httpRequest.version()))
	builder.WriteString(httpResponse.status.String())
	builder.WriteString("\r\n")
	if httpResponse.contentType != "" {
		builder.WriteString(fmt.Sprintf("%s: %s", CONTENT_TYPE, httpResponse.contentType))
	}

	if httpResponse.contentLength > 0 {
		builder.WriteString("\r\n")
		builder.WriteString(fmt.Sprintf("%s: %d", CONTENT_LENGTH, httpResponse.contentLength))
	}

	for key, val := range httpResponse.headersMap {
		builder.WriteString("\r\n")
		builder.WriteString(fmt.Sprintf("%s: %s", key, val))
	}

	builder.WriteString("\r\n")
	builder.WriteString("\r\n")

	builder.WriteString(httpResponse.body)

	newVar := builder.String()
	return newVar
}

func enrichHeaders(req HttpRequest, res HttpResponse) {
	if res.contentType != "" {
		//res.addHeader()
	}
}
