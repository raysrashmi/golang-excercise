package atbashcipher

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var result []string
	var current strings.Builder

	s = strings.ToLower(s)

	for _, ch := range s {
		if unicode.IsLetter(ch) {
			// atbash: a->z, b->y
			encoded := 'z' - (ch - 'a')
			current.WriteRune(encoded)
		} else if unicode.IsDigit(ch) {
			current.WriteRune(ch)
		} else {
			continue // ignore punctuation
		}

		// group every 5 chars
		if current.Len() == 5 {
			result = append(result, current.String())
			current.Reset()
		}
	}

	// leftover
	if current.Len() > 0 {
		result = append(result, current.String())
	}

	return strings.Join(result, " ")
}