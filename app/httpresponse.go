package main

type HttpResponse struct {
	status        HttpStatus
	contentType   string
	contentLength int
	body          string
}

func CreateHttpResponse(httpStatus HttpStatus, contentType string, body string) HttpResponse {
	return HttpResponse{
		status:        httpStatus,
		contentType:   contentType,
		contentLength: len(body),
		body:          body,
	}
}

var NotFoundResponse = CreateHttpResponse(StatusNotFound, "", "")
var RootResponse = CreateHttpResponse(StatusOk, "", "")

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
