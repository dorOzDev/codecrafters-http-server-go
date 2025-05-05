package main

import "fmt"

type HttpStatus struct {
	code  int
	value string
}

var StatusOk = HttpStatus{code: 200, value: "OK"}
var StatusNotFound = HttpStatus{code: 404, value: "Not Found"}

func (s HttpStatus) String() string {
	return fmt.Sprintf("%d %s", s.code, s.value)
}
