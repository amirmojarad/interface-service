package sentences

import "strings"

type sentence struct {
	RawSentence string
	timeRange   string
	tokens      []string
}

func generateSentence(timeRange, rawSentence string) *sentence {
	tokens := strings.Fields(rawSentence)
	for i, token := range tokens {
		if strings.Contains(token, ",") {
			tokens[i] = strings.Replace(token, ",", "", 1)
		}
	}
	return &sentence{
		timeRange:   timeRange,
		tokens:      tokens,
		RawSentence: rawSentence,
	}
}
