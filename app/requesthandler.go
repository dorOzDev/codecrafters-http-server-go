package main

import "strings"

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

var handlers = []Handler{
	RootHandler{},
	EchoHandler{},
	UserAgentHandler{},
	FilesHandler{},
	NotFoundHandler{},
}

var compressList = []Compresser{
	GzipCompresser{},
}

func HandleHttpRequest(request HttpRequest) HttpResponse {
	var resp HttpResponse
	for _, handler := range handlers {
		if handler.accept(request) {
			resp = handler.handleRequest(request)
			resp.enrichHeaders(request)
			break
		}
	}

	compressorName, exists := resp.headersMap[CONTENT_ENCODING]
	if exists {
		compressor := getCompressor(compressorName)
		compressedBody, err := compressor.compress(resp.body)
		if err != nil {
			panic(err)
		}
		resp.contentLength = len(compressedBody)
		resp.body = compressedBody

	}

	return resp
}

func getCompressor(compressorName string) Compresser {
	var res Compresser

	for _, compressor := range compressList {
		if strings.EqualFold(compressor.compressorName(), compressorName) {
			res = compressor
		}
	}

	return res
}
