package main

import "strings"

type UserAgentHandler struct{}

func (e UserAgentHandler) accept(httpRequest HttpRequest) bool {
	return strings.HasPrefix(httpRequest.Path(), "/user-agent")
}

func (e UserAgentHandler) handleRequest(httpRequest HttpRequest) HttpResponse {
	return CreateHttpResponse(StatusOk, "text/plain", strings.TrimPrefix(httpRequest.Path(), "/user-agent"))
}
