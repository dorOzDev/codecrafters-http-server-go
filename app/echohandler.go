package main

import "strings"

type EchoHandler struct{}

func (e EchoHandler) accept(httpRequest HttpRequest) bool {
	return strings.HasPrefix(httpRequest.Path(), "/echo/")
}

func (e EchoHandler) handleRequest(httpRequest HttpRequest) HttpResponse {
	return CreateHttpResponse(StatusOk, "text/plain", strings.TrimPrefix(httpRequest.Path(), "/echo/"))
}
