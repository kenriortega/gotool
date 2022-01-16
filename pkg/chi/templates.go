package chi

var (
	// Init files

	gomodTemplate = `module {{ .ModuleName }}

go 1.17`
	gitignoreTemplate = `.vscode/
.env`

	// CMD folder

	mainTemplate = `// Code generated by go.
// Generated {{ .Timestamp }}
package main

import (
	"runtime"
	"log"
	"github.com/joho/godotenv"
	"{{ .ModuleName }}/cmd/api"

)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)

}
func main() {
	api.Start()
}`

	routesTemplate = `// Code generated by go.
// Generated {{ .Timestamp }}
package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"{{ .ModuleName }}/internal/handlers"
)

func Routes(router *chi.Mux, handler *handlers.BasicHandler) {
	router.Get("/", handler.Index)
	router.Get("/metrics", promhttp.Handler().ServeHTTP)
}`

	startTemplate = `// Code generated by go.
// Generated {{ .Timestamp }}
package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"{{ .ModuleName }}/internal/handlers"
)

func Start() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	h := handlers.New()

	Routes(r, h)
	log.Println("Data importer API starting...")
	log.Fatal(http.ListenAndServe(":3000", r))

}`

	// Internals

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

import (
	"net/http"

	"{{ .ModuleName }}/internal/ports"
)

type BasicHandler struct {
}
func New() *BasicHandler {
	return &BasicHandler{}
}

//This line is for get feedback in case we are not implementing the interface correctly
var _ ports.IBasicHandler = (*BasicHandler)(nil)

func (b *BasicHandler) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}`

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
