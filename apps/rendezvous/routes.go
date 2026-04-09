package rendezvous

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/sessions", createSessionHandler)
	mux.HandleFunc("/session", JoinSessionHandler)
}
