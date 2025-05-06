package main

type HttpRequest interface {
	method() string
	path() string
	raw() string
	version() string
	headers() map[string]string
	queryParams() map[string]string
}

type GetRequest struct {
	methodValue    string
	pathValue      string
	versionValue   string
	rawValue       string
	headersMap     map[string]string
	queryParamsMap map[string]string
}

func (r GetRequest) method() string                 { return r.methodValue }
func (r GetRequest) path() string                   { return r.pathValue }
func (r GetRequest) raw() string                    { return r.rawValue }
func (r GetRequest) version() string                { return r.versionValue }
func (r GetRequest) headers() map[string]string     { return r.headersMap }
func (r GetRequest) queryParams() map[string]string { return r.queryParamsMap }

type PostRequest struct {
	methodValue    string
	pathValue      string
	versionValue   string
	rawValue       string
	headersMap     map[string]string
	queryParamsMap map[string]string
	bodyValue      string
}

func (r PostRequest) method() string                 { return r.methodValue }
func (r PostRequest) path() string                   { return r.pathValue }
func (r PostRequest) raw() string                    { return r.rawValue }
func (r PostRequest) version() string                { return r.versionValue }
func (r PostRequest) headers() map[string]string     { return r.headersMap }
func (r PostRequest) queryParams() map[string]string { return r.queryParamsMap }

func (r PostRequest) body() string {
	return r.bodyValue
}

const (
	GET  = "GET"
	POST = "POST"
)

type HttpMethod struct{}

func (HttpMethod) get() string  { return GET }
func (HttpMethod) post() string { return POST }
