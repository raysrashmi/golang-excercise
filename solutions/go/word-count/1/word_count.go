package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := make(Frequency)

	// case-insensitive (optional depending on requirement)
	phrase = strings.ToLower(phrase)

	re := regexp.MustCompile(`[a-z0-9]+(?:'[a-z0-9]+)*`)
    
	words := re.FindAllString(phrase, -1)

	for _, word := range words {
		freq[word]++
	}

	return freq
}