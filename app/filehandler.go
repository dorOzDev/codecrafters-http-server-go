package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const filesHandlerPath = "/files/"
const directoryFlag = "--directory"

type FilesHandler struct{}

func (e FilesHandler) accept(httpRequest HttpRequest) bool {
	return strings.HasPrefix(httpRequest.path(), filesHandlerPath)
}

func (e FilesHandler) handleRequest(httpRequest HttpRequest) HttpResponse {
	dir, err := getFlagValue(directoryFlag)

	if err != nil {
		panic(err)
	}

	fileName := strings.TrimPrefix(httpRequest.path(), filesHandlerPath)
	absolutePath := filepath.Join(dir, fileName)
	fmt.Print("looking for file: ", absolutePath)

	data, err := os.ReadFile(absolutePath)
	if err != nil {
		return NotFoundResponse
	}

	return CreateHttpResponse(StatusOk, ContentType{}.octet(), string(data))
}
