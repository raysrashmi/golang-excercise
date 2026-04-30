package parallelletterfrequency

import (
	"strings"
	"sync"
	"unicode"
)

type FreqMap map[rune]int

func Frequency(text string) FreqMap {
	freqMap := FreqMap{}
	for _, r := range strings.ToLower(text) {
		if unicode.IsLetter(r) {
			freqMap[r]++
		}
	}
	return freqMap
}

func ConcurrentFrequency(texts []string) FreqMap {
	ch := make(chan FreqMap, len(texts))
	var wg sync.WaitGroup

	for _, text := range texts {
		wg.Go(func() {
			ch <- Frequency(text)
		})
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	result := FreqMap{}
	for partial := range ch {
		for r, count := range partial {
			result[r] += count
		}
	}
	return result
}