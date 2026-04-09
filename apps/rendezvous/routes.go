package rendezvous

import "net/http"

type Handler struct {
	store *SessionStore
}

func NewHandler(store *SessionStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", h.healthHandler)
	mux.HandleFunc("/sessions", h.createSessionHandler)
	mux.HandleFunc("/sessions/join", h.joinSessionHandler)
}
