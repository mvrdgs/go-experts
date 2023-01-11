package handlers

import (
	"net/http"
	"net/http/httptest"
)

func ExecuteRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}
