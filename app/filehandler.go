package main

import (
	"os"
	"strings"
)

const filesHandlerPath = "/files/"

type FilesHandler struct{}

func (e FilesHandler) accept(httpRequest HttpRequest) bool {
	return strings.HasPrefix(httpRequest.path(), filesHandlerPath)
}

func (e FilesHandler) handleRequest(httpRequest HttpRequest) HttpResponse {
	fileName := strings.TrimPrefix(httpRequest.path(), filesHandlerPath)
	data, err := os.ReadFile(fileName)
	if err != nil {
		return NotFoundResponse
	}

	return CreateHttpResponse(StatusOk, ContentType{}.octet(), string(data))
}
