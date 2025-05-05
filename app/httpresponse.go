package main

type HttpResponse struct {
	StatusCode    int
	ContentType   string
	ContentLength int
	Body          string
}

func CreateHttpResponse(status int, contentType string, body string) HttpResponse {
	return HttpResponse{
		StatusCode:    status,
		ContentType:   contentType,
		ContentLength: len(body),
		Body:          body,
	}
}

var NotFoundResponse = CreateHttpResponse(404, "", "")
var RootResponse = CreateHttpResponse(200, "", "")
