package rendezvous

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

type JoinSessionResponse struct {
	SessionID string `json:"session_id"`
	Status    string `json:"status"`
}

type HostCredentials struct {
	SessionID string
	HostToken string
}

type Session struct {
	SessionID string
	JoinCode  string
	HostToken string
	Status    string
}
