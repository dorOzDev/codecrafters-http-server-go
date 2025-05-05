package main

type HttpRequest interface {
	method() string
	path() string
	raw() string
	version() string
	headers() map[string]string
}

type GetRequest struct {
	methodValue  string
	pathValue    string
	versionValue string
	rawValue     string
	headersMap   map[string]string
	queryParams  map[string]string
}

func (r GetRequest) method() string             { return r.methodValue }
func (r GetRequest) path() string               { return r.pathValue }
func (r GetRequest) raw() string                { return r.rawValue }
func (r GetRequest) version() string            { return r.versionValue }
func (r GetRequest) headers() map[string]string { return r.headersMap }
