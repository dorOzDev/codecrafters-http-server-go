package main

import "strings"

type UserAgentHandler struct{}

func (e UserAgentHandler) accept(httpRequest HttpRequest) bool {
	return strings.HasPrefix(httpRequest.path(), "/user-agent")
}

func (e UserAgentHandler) handleRequest(httpRequest HttpRequest) HttpResponse {
	userAgent := httpRequest.headers()["User-Agent"]
	return CreateHttpResponse(StatusOk, ContentType{}.text(), []byte(userAgent))
}
