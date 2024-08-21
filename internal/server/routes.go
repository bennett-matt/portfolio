package server

import (
	"encoding/json"
	"log"
	"net/http"

	"portfolio/cmd/web"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/health", s.healthHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	r.Handler(http.MethodGet, "/assets/*filepath", fileServer)

	dynamic := alice.New(s.sessionManager.LoadAndSave, noSurf)
	r.Handler(http.MethodGet, "/", dynamic.ThenFunc(s.home))
	r.Handler(http.MethodGet, "/about", dynamic.ThenFunc(s.about))
	r.Handler(http.MethodGet, "/projects", dynamic.ThenFunc(s.projects))
	r.Handler(http.MethodGet, "/contact", dynamic.ThenFunc(s.contact))
	r.Handler(http.MethodGet, "/sign-in", dynamic.ThenFunc(s.signIn))
	r.Handler(http.MethodGet, "/sign-up", dynamic.ThenFunc(s.signUp))

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
