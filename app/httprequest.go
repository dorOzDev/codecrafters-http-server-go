package main

import (
	"strings"
	"sync"
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
	methodValue    string
	pathValue      string
	versionValue   string
	rawValue       string
	headersMap     map[string]string
	queryParamsMap map[string]string
}

func (r GetRequest) method() string                         { return r.methodValue }
func (r GetRequest) path() string                           { return r.pathValue }
func (r GetRequest) raw() string                            { return r.rawValue }
func (r GetRequest) version() string                        { return r.versionValue }
func (r GetRequest) headers() map[string]string             { return r.headersMap }
func (r GetRequest) queryParams() map[string]string         { return r.queryParamsMap }
func (r GetRequest) hasHeader(header string) (string, bool) { return hasHeader(r, header) }

type PostRequest struct {
	methodValue    string
	pathValue      string
	versionValue   string
	rawValue       string
	headersMap     map[string]string
	queryParamsMap map[string]string
	bodyValue      string
}

func (r PostRequest) method() string                         { return r.methodValue }
func (r PostRequest) path() string                           { return r.pathValue }
func (r PostRequest) raw() string                            { return r.rawValue }
func (r PostRequest) version() string                        { return r.versionValue }
func (r PostRequest) headers() map[string]string             { return r.headersMap }
func (r PostRequest) queryParams() map[string]string         { return r.queryParamsMap }
func (r PostRequest) hasHeader(header string) (string, bool) { return hasHeader(r, header) }

func (r PostRequest) body() string {
	return r.bodyValue
}

const (
	GET  = "GET"
	POST = "POST"
)

const (
	ACCEPT_ENCODING  = "Accept-Encoding"
	CONTENT_ENCODING = "Content-Encoding"
)

func hasHeader(httpRequest HttpRequest, header string) (string, bool) {
	if httpRequest.headers() == nil {
		return "", false
	}

	val, exists := httpRequest.headers()[header]
	return val, exists
}

var (
	supportedEncodingMap map[string]struct{}
	once                 sync.Once
)

func isSupportedEncoding(encodingArray []string) (string, bool) {
	once.Do(func() {
		supportedEncodingMap = map[string]struct{}{
			"gzip": {},
		}
	})

	for _, encoding := range encodingArray {
		_, exists := supportedEncodingMap[strings.ToLower(encoding)]
		if exists {
			return encoding, true
		}
	}

	return "", false
}

func parseAcceptEncoding(headerValue string) []string {
	encodings := strings.Split(headerValue, ",")
	for i := range encodings {
		encodings[i] = strings.TrimSpace(encodings[i])
	}
	return encodings
}
