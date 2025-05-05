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
