package server

import (
	"net/http"
	"portfolio/cmd/web"
	"portfolio/cmd/web/layouts"
	"portfolio/internal/models"

	"github.com/a-h/templ"
	"github.com/justinas/nosurf"
)

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	s.render(w, r, web.Home(), layouts.Base)
}

func (s *Server) about(w http.ResponseWriter, r *http.Request) {
	s.render(w, r, web.About(), layouts.Base)
}

func (s *Server) projects(w http.ResponseWriter, r *http.Request) {
	s.render(w, r, web.Projects(), layouts.Base)
}

func (s *Server) contact(w http.ResponseWriter, r *http.Request) {
	s.render(w, r, web.ContactMe(), layouts.Base)
}

func (s *Server) signIn(w http.ResponseWriter, r *http.Request) {
	s.render(w, r, web.SignIn(), layouts.Auth)
}

func (s *Server) signUp(w http.ResponseWriter, r *http.Request) {
	s.render(w, r, web.SignUp(), layouts.Auth)
}

func (s *Server) render(w http.ResponseWriter, r *http.Request, component templ.Component, layout func(templ.Component, *models.ViewData) templ.Component) {
	data := s.newViewData(r)
	layout(component, data).Render(r.Context(), w)
}

func (s *Server) newViewData(r *http.Request) *models.ViewData {
	return &models.ViewData{
		Flash:           s.sessionManager.PopString(r.Context(), "flash"),
		CSRFToken:       nosurf.Token(r),
		IsAuthenticated: s.isAuthenticated(r),
	}
}
