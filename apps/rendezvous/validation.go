package rendezvous

import (
	"encoding/base64"
	"fmt"
)

func ValidatePublicKey(s string) ([]byte, error) {
	if s == "" {
		return nil, fmt.Errorf("public key is required")
	}

	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("invalid base64 public key: %w", err)
	}

	if len(b) != 32 {
		return nil, fmt.Errorf("invalid X25519 public key length: got %d, want 32", len(b))
	}

	return b, nil
}
