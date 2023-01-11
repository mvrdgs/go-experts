package handlers

import "net/http"

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte("route does not exist"))
}
