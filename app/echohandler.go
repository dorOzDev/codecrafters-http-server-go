package main

import "strings"

const echoHandlerPath = "/echo/"

type EchoHandler struct{}

func (e EchoHandler) accept(httpRequest HttpRequest) bool {
	return strings.HasPrefix(httpRequest.path(), echoHandlerPath)
}

func (e EchoHandler) handleRequest(httpRequest HttpRequest) HttpResponse {
	val, exists := httpRequest.hasHeader(ACCEPT_ENCODING)
	resp := CreateHttpResponse(StatusOk, ContentType{}.text(), strings.TrimPrefix(httpRequest.path(), echoHandlerPath))
	if exists {
		strings.Split(val, ",")
		headerValue, exists := isSupportedEncoding(parseAcceptEncoding(val))
		if exists {
			resp.addHeader(CONTENT_ENCODING, headerValue)
		} else {

		}
	}
	return resp
}
