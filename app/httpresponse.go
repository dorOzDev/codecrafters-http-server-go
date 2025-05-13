package main

import (
	"fmt"
	"strconv"
	"strings"
)

type HttpResponse struct {
	status        HttpStatus
	contentType   string
	contentLength int
	body          []byte
	headersMap    map[string]string
}

func CreateHttpResponse(httpStatus HttpStatus, contentType string, body []byte) HttpResponse {
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

var NotFoundResponse = CreateHttpResponse(StatusNotFound, "", []byte{})
var RootResponse = CreateHttpResponse(StatusOk, "", []byte{})
var CreatedResponse = CreateHttpResponse(StatusCreated, "", []byte{})
var UnexpectedError = CreateHttpResponse(StatusInternalError, "", []byte{})

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

	for key, val := range resp.headersMap {
		builder.WriteString("\r\n")
		builder.WriteString(fmt.Sprintf("%s: %s", key, val))
	}

	builder.WriteString("\r\n")
	builder.WriteString("\r\n")

	builder.WriteString(string(resp.body))

	newVar := builder.String()
	return newVar
}

func (resp *HttpResponse) enrichHeaders(req HttpRequest) {
	if resp.contentType != "" {
		resp.AddHeader(CONTENT_TYPE, resp.contentType)
	}

	if resp.contentLength > 0 {
		resp.AddHeader(CONTENT_LENGTH, strconv.Itoa(resp.contentLength))
	}
	val, exists := req.hasHeader(ACCEPT_ENCODING)
	if exists {
		strings.Split(val, ",")
		headerValue, exists := isSupportedEncoding(parseAcceptEncoding(val))
		if exists {
			resp.AddHeader(CONTENT_ENCODING, headerValue)
		}
	}
	val, exists = req.hasHeader(CONNECTION)
	if exists && strings.ToLower(val) == "close" {
		resp.AddHeader(CONNECTION, val)
	}
}
