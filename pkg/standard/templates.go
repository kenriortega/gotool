package standard

var (
	gomodTemplate = `module {{ .ModuleName }}

go 1.17`
	gitignoreTemplate = `.vscode/`

	mainTemplate = `// Code generated by go.
// Generated {{ .Timestamp }}
package main

import "fmt"

func main() {
	fmt.Printf("%d\n", "Hello world")
}`
	domainTemplate = `// Code generated by go.
// Generated {{ .Timestamp }}
package domain

type Basic struct{}`
	portsTemplate = `// Code generated by go.
// Generated {{ .Timestamp }}
package ports

type IBasicHandler interface{}
type IBasicService interface{}`

	handlerTemplate = `// Code generated by go.
// Generated {{ .Timestamp }}
package handlers

import "{{ .ModuleName }}/internal/ports"

type BasicHandler struct {
}
func New() *BasicHandler {
	return &BasicHandler{}
}

//This line is for get feedback in case we are not implementing the interface correctly
var _ ports.IBasicHandler = (*BasicHandler)(nil)`

	serviceTemplate = `// Code generated by go.
// Generated {{ .Timestamp }}
package services

import "{{ .ModuleName }}/internal/ports"

type BasicService struct {
}
func New() *BasicService {
	return &BasicService{}
}

//This line is for get feedback in case we are not implementing the interface correctly
var _ ports.IBasicService = (*BasicService)(nil)`
)
