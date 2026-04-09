package rendezvous

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{"ok"})
}

func createSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	var req CreateSessionRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&req); err != nil {
		log.Printf("invalid request body: %v", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.ProtocolVersion != 1 {
		log.Printf("req protocol version %d differs from server version", req.ProtocolVersion)
		http.Error(w, "invalid params", http.StatusBadRequest)
		return
	}

	session, _, err := CreateSession(req.PublicKey, req.ClientName)
	if err != nil {
		if errors.Is(err, ErrInvalidPublicKey) {
			http.Error(w, "invalid public key encoding", http.StatusBadRequest)
			return
		}
		log.Printf("failed to create session: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	resp := CreateSessionResponse{
		SessionID: session.SessionID,
		JoinCode:  session.JoinCode,
		HostToken: session.HostToken,
		Status:    session.Status,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func JoinSessionHandler(w http.ResponseWriter, r *http.Request) {
}
