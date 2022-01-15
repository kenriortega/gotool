package standard

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
	"time"
)

var (
	gomodTemplate = `module {{ .ModuleName }}

go 1.17`
	gitignoreTemplate = `.vscode/`
	mainTemplate      = `// Code generated by go.
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

import "{{ .ModuleName }}/ports"

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

import "{{ .ModuleName }}/ports"

type BasicService struct {
}
func New() *BasicService {
	return &BasicService{}
}

//This line is for get feedback in case we are not implementing the interface correctly
var _ ports.IBasicService = (*BasicService)(nil)`
)

func FoldersStructure(out, name string) {
	folders := []string{
		fmt.Sprintf("%s/%s/cmd", out, name),
		fmt.Sprintf("%s/%s/pkg", out, name),
		fmt.Sprintf("%s/%s/internal/domain", out, name),
		fmt.Sprintf("%s/%s/internal/ports", out, name),
		fmt.Sprintf("%s/%s/internal/handlers", out, name),
		fmt.Sprintf("%s/%s/internal/services", out, name),
	}

	for _, folder := range folders {
		if err := os.MkdirAll(folder, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		// Boilerplate files
		archives := strings.Split(folder, "/")
		archive := archives[len(archives)-1:][0]
		switch archive {
		case "cmd":
			file, err := os.Create(fmt.Sprintf("%s/main.go", folder))
			if err != nil {
				fmt.Errorf("%v", err)
				os.Exit(1)
			}
			defer file.Close()

			template.Must(
				template.New("").Parse(mainTemplate)).
				Execute(file, struct {
					Timestamp time.Time
				}{
					time.Now(),
				})
		case "domain":
			file, err := os.Create(fmt.Sprintf("%s/basic.go", folder))
			if err != nil {
				fmt.Errorf("%v", err)
				os.Exit(1)
			}
			defer file.Close()

			template.Must(
				template.New("").Parse(domainTemplate)).
				Execute(file, struct {
					Timestamp time.Time
				}{
					time.Now(),
				})
		case "ports":
			file, err := os.Create(fmt.Sprintf("%s/basic.go", folder))
			if err != nil {
				fmt.Errorf("%v", err)
				os.Exit(1)
			}
			defer file.Close()

			template.Must(
				template.New("").Parse(portsTemplate)).
				Execute(file, struct {
					Timestamp time.Time
				}{
					time.Now(),
				})
		case "handlers":
			file, err := os.Create(fmt.Sprintf("%s/basic.go", folder))
			if err != nil {
				fmt.Errorf("%v", err)
				os.Exit(1)
			}
			defer file.Close()

			template.Must(
				template.New("").Parse(handlerTemplate)).
				Execute(file, struct {
					Timestamp  time.Time
					ModuleName string
				}{
					time.Now(),
					name,
				})
		case "services":
			file, err := os.Create(fmt.Sprintf("%s/basic.go", folder))
			if err != nil {
				fmt.Errorf("%v", err)
				os.Exit(1)
			}
			defer file.Close()

			template.Must(
				template.New("").Parse(serviceTemplate)).
				Execute(file, struct {
					Timestamp  time.Time
					ModuleName string
				}{
					time.Now(),
					name,
				})
		}
	}
}
func FilesInit(path, name string, files ...string) {
	for _, fileInit := range files {

		switch fileInit {
		case "go.mod":
			file, err := os.Create(fmt.Sprintf("%s/%s/%s", path, name, fileInit))
			if err != nil {
				fmt.Errorf("%v", err)
				os.Exit(1)
			}
			defer file.Close()

			template.Must(
				template.New("").Parse(gomodTemplate)).
				Execute(file, struct {
					Timestamp  time.Time
					ModuleName string
				}{
					time.Now(),
					name,
				})
		case ".gitignore":
			file, err := os.Create(fmt.Sprintf("%s/%s/%s", path, name, fileInit))
			if err != nil {
				fmt.Errorf("%v", err)
				os.Exit(1)
			}
			defer file.Close()

			template.Must(
				template.New("").Parse(gitignoreTemplate)).
				Execute(file, struct {
					Timestamp time.Time
				}{
					time.Now(),
				},
				)
		}
	}
}
