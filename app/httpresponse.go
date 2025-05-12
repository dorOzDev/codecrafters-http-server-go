package main

import (
	"fmt"
	"strings"
)

type HttpResponse struct {
	status        HttpStatus
	contentType   string
	contentLength int
	body          string
	headersMap    map[string]string
}

func CreateHttpResponse(httpStatus HttpStatus, contentType string, body string) HttpResponse {
	return HttpResponse{
		status:        httpStatus,
		contentType:   contentType,
		contentLength: len(body),
		body:          body,
		headersMap:    make(map[string]string),
	}
}

func (r HttpResponse) AddHeader(headerName string, headerValue string) {
	r.headersMap[headerName] = headerValue
}

var NotFoundResponse = CreateHttpResponse(StatusNotFound, "", "")
var RootResponse = CreateHttpResponse(StatusOk, "", "")
var CreatedResponse = CreateHttpResponse(StatusCreated, "", "")
var UnexpectedError = CreateHttpResponse(StatusInternalError, "", "")

const (
	JSON  = "application/json"
	XML   = "application/xml"
	HTML  = "text/html"
	TEXT  = "text/plain"
	JPEG  = "image/jpeg"
	PNG   = "image/png"
	OCTET = "application/octet-stream"
)

// Optional: Use a struct for organization
type ContentType struct{}

func (ContentType) json() string  { return JSON }
func (ContentType) xml() string   { return XML }
func (ContentType) html() string  { return HTML }
func (ContentType) text() string  { return TEXT }
func (ContentType) jpeg() string  { return JPEG }
func (ContentType) png() string   { return PNG }
func (ContentType) octet() string { return OCTET }

func (resp HttpResponse) reformatResponse(req HttpRequest) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("%s ", req.version()))
	builder.WriteString(resp.status.String())
	builder.WriteString("\r\n")
	if resp.contentType != "" {
		builder.WriteString(fmt.Sprintf("%s: %s", CONTENT_TYPE, resp.contentType))
	}

	if resp.contentLength > 0 {
		builder.WriteString("\r\n")
		builder.WriteString(fmt.Sprintf("%s: %d", CONTENT_LENGTH, resp.contentLength))
	}

	for key, val := range resp.headersMap {
		builder.WriteString("\r\n")
		builder.WriteString(fmt.Sprintf("%s: %s", key, val))
	}

	builder.WriteString("\r\n")
	builder.WriteString("\r\n")

	builder.WriteString(resp.body)

	newVar := builder.String()
	return newVar
}
