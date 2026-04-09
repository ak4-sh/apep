package rendezvous

import (
	_ "embed"
	"strings"
)

//go:embed wordlist.txt
var wordListRaw string

var wordList []string

func init() {
	for line := range strings.SplitSeq(wordListRaw, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		wordList = append(wordList, line)
	}
}
