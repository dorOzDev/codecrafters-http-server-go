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

	switch concrete := httpRequest.(type) {
	case *GetRequest:
		return handleGetRequest(concrete)
	case *PostRequest:
		return handlePostRequest(concrete)
	default:
		return UnexpectedError
	}
}

func handlePostRequest(request *PostRequest) HttpResponse {
	absolutePath, err := getAbsolutePath(request)
	if err != nil {
		fmt.Print("unable to get absolute path: ", err)
		return UnexpectedError
	}

	fmt.Print("writing to file: ", absolutePath)
	err = os.WriteFile(absolutePath, []byte(request.bodyValue), 0644)
	if err != nil {
		fmt.Print("unable to write to file: ", err)
		return UnexpectedError
	}
	fmt.Print("successfully written to file: ", absolutePath)
	return CreatedResponse
}

func handleGetRequest(request *GetRequest) HttpResponse {
	absolutePath, err := getAbsolutePath(request)
	if err != nil {
		fmt.Print("unable to get absolute path: ", err)
		return UnexpectedError
	}

	fmt.Print("looking for file: ", absolutePath)
	data, err := os.ReadFile(absolutePath)
	if err != nil {
		return NotFoundResponse
	}

	return CreateHttpResponse(StatusOk, ContentType{}.octet(), data)
}

func getAbsolutePath(request HttpRequest) (string, error) {
	dir, err := getFlagValue(directoryFlag)
	if err != nil {
		return "", err
	}

	fileName := strings.TrimPrefix(request.path(), filesHandlerPath)
	absolutePath := filepath.Join(dir, fileName)
	return absolutePath, nil
}
