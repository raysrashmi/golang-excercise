package anagram

import "fmt"
import "strings"

func Detect(subject string, candidates []string) []string {
    a := []string{}
    subjectL := strings.ToLower(subject)
	for _ , str := range candidates {
        if isSubset(subjectL, strings.ToLower(str)) {
            a = append(a, str)
        }
    }
    fmt.Println(a)
    return a
}

func isSubset(na, nb string) bool {    
	// A word is not an anagram of itself
	if string(na) == string(nb) {
		return false
	}

	if len(na) != len(nb) {
		return false
	}

	count := make(map[rune]int)

	for _, ch := range na {
		count[ch]++
	}

	for _, ch := range nb {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}

	return true
}
