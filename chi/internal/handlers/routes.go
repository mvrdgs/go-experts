package handlers

import (
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) MountHandlers() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	s.Router.Get("/", HelloWorld)
	s.Router.NotFound(NotFoundHandler)
}
