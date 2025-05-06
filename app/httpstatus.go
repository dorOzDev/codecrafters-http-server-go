package main

import "fmt"

type HttpStatus struct {
	code  int
	value string
}

var StatusOk = HttpStatus{code: 200, value: "OK"}
var StatusNotFound = HttpStatus{code: 404, value: "Not Found"}
var StatusInternalError = HttpStatus{code: 500, value: "Internal Server Error"}
var StatusCreated = HttpStatus{code: 201, value: "Created"}

func (s HttpStatus) String() string {
	return fmt.Sprintf("%d %s", s.code, s.value)
}
