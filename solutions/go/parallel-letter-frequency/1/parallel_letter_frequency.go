package parallelletterfrequency

import (
	"sync"
	"unicode"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string, wg *sync.WaitGroup, ch chan FreqMap) {
	defer wg.Done()
	freqMap := FreqMap{}
	for _, r := range text {
		if unicode.IsLetter(r) {
			freqMap[unicode.ToLower(r)]++
		}
	}
	ch <- freqMap
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	ch := make(chan FreqMap, len(texts))
	var wg sync.WaitGroup

	for _, text := range texts {
		wg.Add(1)
		go Frequency(text, &wg, ch)
	}

	wg.Wait()
	close(ch)

	result := FreqMap{}
	for partial := range ch {
		for r, count := range partial {
			result[r] += count
		}
	}
	return result
}