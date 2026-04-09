// Package rendezvous contains API implementation for Apep's rendezvous server
package rendezvous

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
)

func newSessionID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func newToken(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func newJoinCode() (string, error) {
	if len(wordList) == 0 {
		return "", fmt.Errorf("word list is empty")
	}

	var sb strings.Builder

	for i := range 3 {
		randIdx, err := rand.Int(rand.Reader, big.NewInt(int64(len(wordList))))
		if err != nil {
			return "", err
		}

		if i > 0 {
			sb.WriteByte('-')
		}
		sb.WriteString(wordList[int(randIdx.Int64())])
	}

	randVal, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return "", err
	}

	fmt.Fprintf(&sb, "-%02d", randVal.Int64())

	return sb.String(), nil
}
