package rendezvous

import (
	"sync"
	"time"
)

const (
	sessionTTL      = 30 * time.Minute
	cleanupInterval = 2 * time.Minute
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
	createdAt time.Time
}

func (s Session) Expired() bool {
	return time.Since(s.createdAt) > sessionTTL
}

type SessionStore struct {
	mu         sync.RWMutex
	byID       map[string]*Session
	byJoinCode map[string]*Session
}

func (s *SessionStore) GetByID(id string) (*Session, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	ret, ok := s.byID[id]

	if !ok {
		return nil, false
	}

	if ret.Expired() {
		delete(s.byID, id)
		delete(s.byJoinCode, ret.JoinCode)
		return nil, false
	}
	return ret, true
}

func (s *SessionStore) GetByJoinCode(joinCode string) (*Session, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	ret, ok := s.byJoinCode[joinCode]

	if !ok {
		return nil, false
	}

	if ret.Expired() {
		delete(s.byJoinCode, joinCode)
		delete(s.byID, ret.SessionID)
		return nil, false
	}
	return ret, true
}

func (s *SessionStore) startCleanupLoop() {
	ticker := time.NewTicker(cleanupInterval)

	go func() {
		defer ticker.Stop()

		s.cleanup()

		for range ticker.C {
			s.cleanup()
		}
	}()
}

func (s *SessionStore) cleanup() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, session := range s.byID {
		if session.Expired() {
			delete(s.byID, session.SessionID)
			delete(s.byJoinCode, session.JoinCode)
		}
	}
}

func NewSessionStore() *SessionStore {
	store := SessionStore{
		byID:       make(map[string]*Session),
		byJoinCode: make(map[string]*Session),
	}

	store.startCleanupLoop()
	return &store
}

func (s *SessionStore) Put(session *Session) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.byID[session.SessionID] = session
	s.byJoinCode[session.JoinCode] = session
}
