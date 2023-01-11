package main

import (
	"net/http"

	"github.com/mvrdgs/go-experts/test/internal/handlers"
)

func main() {
	s := handlers.CreateNewServer()
	s.MountHandlers()
	http.ListenAndServe(":8080", s.Router)
}
