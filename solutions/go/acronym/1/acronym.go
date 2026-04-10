// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	wordsList := splitBySpaceAndHyphen(s)
    acro := ""
	for _, word := range wordsList {
        word := strings.ReplaceAll(word, "_", "")
        acro = acro + strings.ToUpper(string(word[0]))
    }

    return acro
}

func splitBySpaceAndHyphen(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return unicode.IsSpace(r) || r == '-'
	})
}
