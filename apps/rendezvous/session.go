package rendezvous

import (
	"errors"
	"log"
)

var (
	ErrInvalidProtocolVersion = errors.New("invalid protocol version")
	ErrInvalidPublicKey       = errors.New("invalid public key")
	ErrSessionNotFound        = errors.New("session not found")
)

func CreateSession(publicKey string, clientName string) (*Session, *HostCredentials, error) {
	publicKeyBytes, err := ValidatePublicKey(publicKey)
	if err != nil {
		return nil, nil, ErrInvalidPublicKey
	}

	_ = publicKeyBytes

	sessionID, err := NewSessionID()
	if err != nil {
		log.Printf("failed to generate session id: %v", err)
		return nil, nil, err
	}

	hostToken, err := NewToken(32)
	if err != nil {
		log.Printf("failed to generate host token: %v", err)
		return nil, nil, err
	}

	joinCode, err := NewJoinCode()
	if err != nil {
		log.Printf("failed to generate join code: %v", err)
		return nil, nil, err
	}

	session := &Session{
		SessionID: sessionID,
		JoinCode:  joinCode,
		HostToken: hostToken,
		Status:    "waiting",
	}

	credentials := &HostCredentials{
		SessionID: sessionID,
		HostToken: hostToken,
	}

	return session, credentials, nil
}

func JoinSession(joinCode string, publicKey string) (*Session, error) {
	return nil, ErrSessionNotFound
}
