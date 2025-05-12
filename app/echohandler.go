package main

import "strings"

const echoHandlerPath = "/echo/"

type EchoHandler struct{}

func (e EchoHandler) accept(httpRequest HttpRequest) bool {
	return strings.HasPrefix(httpRequest.path(), echoHandlerPath)
}

func (e EchoHandler) handleRequest(httpRequest HttpRequest) HttpResponse {
	resp := CreateHttpResponse(StatusOk, ContentType{}.text(), strings.TrimPrefix(httpRequest.path(), echoHandlerPath))
	return resp
}
