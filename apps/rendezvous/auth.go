package rendezvous

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
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
	words, err := loadWordList("apps/rendezvous/wordlist.txt")
	if err != nil {
		return "", err
	}
	if len(words) == 0 {
		return "", fmt.Errorf("word list is empty")
	}

	var sb strings.Builder

	for i := range 3 {
		randIdx, err := rand.Int(rand.Reader, big.NewInt(int64(len(words))))
		if err != nil {
			return "", err
		}

		if i > 0 {
			sb.WriteByte('-')
		}
		sb.WriteString(words[int(randIdx.Int64())])
	}

	randVal, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return "", err
	}

	fmt.Fprintf(&sb, "-%02d", randVal.Int64())

	return sb.String(), nil
}

func loadWordList(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	words := make([]string, 0, 7776)
	iter := strings.SplitSeq(string(data), "\n")
	for word := range iter {
		words = append(words, word)
	}

	return words, nil
}
