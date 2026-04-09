package rendezvous

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

type CreateSessionRequest struct {
	ProtocolVersion int    `json:"protocol_version"`
	PublicKey       string `json:"public_key"`
	ClientName      string `json:"client_name,omitempty"`
}

type CreateSessionResponse struct {
	SessionID string `json:"session_id"`
	JoinCode  string `json:"join_code"`
	HostToken string `json:"host_token"`
	Status    string `json:"status"`
}

type JoinSessionRequest struct {
	JoinCode string `json:"join_code"`
}

type JoinSessionResponse struct{}

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

	publicKeyBytes, err := decodeAndValidateX25519PublicKey(req.PublicKey)
	if err != nil {
		log.Printf("public key not matching expected shape: %v", err)
		http.Error(w, "invalid public key encoding", http.StatusBadRequest)
		return
	}

	_ = publicKeyBytes
	_ = req.ClientName

	sessionID, err := newSessionID()
	if err != nil {
		log.Printf("failed to generate session id %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	hostToken, err := newToken(32)
	if err != nil {
		log.Printf("failed to generate host token %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	joinCode, err := newJoinCode()
	if err != nil {
		log.Printf("failed to generate join code %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	resp := CreateSessionResponse{
		SessionID: sessionID,
		JoinCode:  joinCode,
		HostToken: hostToken,
		Status:    "waiting",
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func JoinSessionHandler(w http.ResponseWriter, r *http.Request) {
}
