package main

import (
	"strings"
)

type HttpRequest interface {
	method() string
	path() string
	raw() string
	version() string
	headers() map[string]string
	queryParams() map[string]string
	hasHeader(header string) (string, bool)
}

type GetRequest struct {
	methodValue    httpMethod
	pathValue      string
	versionValue   string
	rawValue       string
	headersMap     map[string]string
	queryParamsMap map[string]string
}

func (r GetRequest) method() string                         { return r.methodValue.value }
func (r GetRequest) path() string                           { return r.pathValue }
func (r GetRequest) raw() string                            { return r.rawValue }
func (r GetRequest) version() string                        { return r.versionValue }
func (r GetRequest) headers() map[string]string             { return r.headersMap }
func (r GetRequest) queryParams() map[string]string         { return r.queryParamsMap }
func (r GetRequest) hasHeader(header string) (string, bool) { return hasHeader(r, header) }

type PostRequest struct {
	methodValue    httpMethod
	pathValue      string
	versionValue   string
	rawValue       string
	headersMap     map[string]string
	queryParamsMap map[string]string
	bodyValue      string
}

func (r PostRequest) method() string                         { return r.methodValue.value }
func (r PostRequest) path() string                           { return r.pathValue }
func (r PostRequest) raw() string                            { return r.rawValue }
func (r PostRequest) version() string                        { return r.versionValue }
func (r PostRequest) headers() map[string]string             { return r.headersMap }
func (r PostRequest) queryParams() map[string]string         { return r.queryParamsMap }
func (r PostRequest) hasHeader(header string) (string, bool) { return hasHeader(r, header) }

func (r PostRequest) body() string {
	return r.bodyValue
}

type httpMethod struct {
	value string
}

var (
	GET         = httpMethod{value: "GET"}
	POST        = httpMethod{value: "POST"}
	UNSUPPORTED = httpMethod{value: "UNSUPPORTED METHOD"}
)

var supportedMethods = []httpMethod{GET, POST}

func GetHttpMethod(value string) (httpMethod, bool) {
	upperCaseValue := strings.ToUpper(value)
	for _, method := range supportedMethods {
		if method.value == upperCaseValue {
			return method, true
		}
	}

	return UNSUPPORTED, false
}

func (actual httpMethod) Equals(expected httpMethod) bool {
	return expected.value == actual.value
}

const (
	ACCEPT_ENCODING  = "Accept-Encoding"
	CONTENT_ENCODING = "Content-Encoding"
	CONTENT_LENGTH   = "Content-Length"
	CONTENT_TYPE     = "Content-Type"
)

func hasHeader(httpRequest HttpRequest, header string) (string, bool) {
	if httpRequest.headers() == nil {
		return "", false
	}

	val, exists := httpRequest.headers()[header]
	return val, exists
}
