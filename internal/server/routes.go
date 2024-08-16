package server

import (
	"encoding/json"
	"log"
	"net/http"

	"portfolio/cmd/web"

	"github.com/a-h/templ"
	"github.com/julienschmidt/httprouter"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := httprouter.New()
	r.HandlerFunc(http.MethodGet, "/", s.HelloWorldHandler)

	r.HandlerFunc(http.MethodGet, "/health", s.healthHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	r.Handler(http.MethodGet, "/assets/*filepath", fileServer)
	r.Handler(http.MethodGet, "/web", templ.Handler(web.HelloForm()))
	r.Handler(http.MethodGet, "/home", templ.Handler(web.Home()))
	r.Handler(http.MethodGet, "/about", templ.Handler(web.About()))
	r.Handler(http.MethodGet, "/projects", templ.Handler(web.Projects()))
	r.Handler(http.MethodGet, "/contact", templ.Handler(web.ContactMe()))
	r.Handler(http.MethodGet, "/sign-in", templ.Handler(web.SignIn()))
	r.Handler(http.MethodGet, "/sign-up", templ.Handler(web.SignUp()))
	r.HandlerFunc(http.MethodPost, "/hello", web.HelloWebHandler)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
