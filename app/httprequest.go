package main

type HttpRequest interface {
	Method() string
	Path() string
	Raw() string
	Version() string
}

type GetRequest struct {
	MethodValue  string
	PathValue    string
	VersionValue string
	RawValue     string
	Headers      map[string]string
	QueryParams  map[string]string
}

func (r GetRequest) Method() string  { return r.MethodValue }
func (r GetRequest) Path() string    { return r.PathValue }
func (r GetRequest) Raw() string     { return r.RawValue }
func (r GetRequest) Version() string { return r.VersionValue }
