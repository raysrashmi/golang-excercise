// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
    "strings"
    "unicode"
    )

// reply of remark
func Hey(remark string) string {
    remark = strings.TrimSpace(remark)
   switch {
case isAllUpper(remark) && strings.HasSuffix(remark, "?"):
	return "Calm down, I know what I'm doing!"

case isAllUpper(remark):
	return "Whoa, chill out!"

case strings.HasSuffix(remark, "?"):
	return "Sure."

case len(strings.ReplaceAll(remark, " ", "")) == 0:
	return "Fine. Be that way!"

default:
	return "Whatever."
}
	
}

func isAllUpper(s string) bool {
	hasLetter := false

	for _, ch := range s {
		if unicode.IsLetter(ch) {
			hasLetter = true
			if !unicode.IsUpper(ch) {
				return false
			}
		}
	}
	return hasLetter
}